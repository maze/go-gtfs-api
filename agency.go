package main

import (
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

type AgencyApi struct {
	C  db.Collection
	DB db.Database
}

func (self *AgencyApi) List() ([]Agency, error) {

	var agencies []Agency

	count, err := self.C.Count(nil)

	agencies = make([]Agency, 0, count)

	q, err := self.C.Query()

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
