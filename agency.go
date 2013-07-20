package main

import (
	"encoding/json"
	esCore "github.com/mattbaird/elastigo/core"
	"menteslibres.net/gosexy/db"
)

type Agency struct {
	AgencyId string `field:"agency_id"`
	Name     string `field:"agency_name"`
	URL      string `field:"agency_url"`
	Timezone string `field:"agency_timezone"`
	Lang     string `field:"agency_lang"`
	Phone    string `field:"agency_phone"`
}

const AgencyCollection = `agency`
const AgencySearchIndex = `gtfs_agency`

type AgencyApi struct {
	C  db.Collection
	DB db.Database
}

func (self *AgencyApi) Search(query string) ([]Agency, error) {

	req := map[string]interface{}{
		"query": map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"agency_name": map[string]interface{}{
					"query":          query,
					"operator":       "and",
					"max_expansions": 8,
				},
			},
		},
	}

	res, err := esCore.SearchRequest(true, AgencySearchIndex, "", req, "")

	if err != nil {
		return nil, err
	}

	conds := make(db.Or, 0, res.Hits.Total)

	for i, _ := range res.Hits.Hits {
		item := struct {
			AgencyId string `json:"agency_id"`
		}{}
		json.Unmarshal(res.Hits.Hits[i].Source, &item)
		conds = append(conds, db.Cond{"agency_id": item.AgencyId})
	}

	return self.query(conds)
}

func (self *AgencyApi) List() ([]Agency, error) {
	return self.query(nil)
}

func (self *AgencyApi) Get(AgencyId string) (*Agency, error) {

	q, err := self.C.Query(db.Cond{
		"agency_id": AgencyId,
	})

	if err != nil {
		return nil, err
	}

	agency := Agency{}

	err = q.One(&agency)
	if err != nil {
		return nil, err
	}

	return &agency, nil
}

func (self *AgencyApi) query(conds ...interface{}) ([]Agency, error) {

	var agencies []Agency

	count, err := self.C.Count(conds...)

	agencies = make([]Agency, 0, count)

	q, err := self.C.Query(conds...)

	if err != nil {
		return nil, err
	}

	for {
		agency := Agency{}
		err := q.Next(&agency)
		if err != nil {
			if err != db.ErrNoMoreRows {
				return nil, err
			}
			break
		}
		agencies = append(agencies, agency)
	}

	return agencies, nil
}
