package main

import (
	"menteslibres.net/gosexy/db"
)

type Frequencies struct {
	TripId      string `field:"service_id"`
	StartTime   string `field:"start_time"`
	EndTime     string `field:"end_time"`
	HeadwaySecs int    `field:"headway_secs"`
	ExactTimes  string `field:"exact_times"`
}

const FrequenciesCollection = `frequencies`

type FrequenciesApi struct {
	C  db.Collection
	DB db.Database
}

func (self *FrequenciesApi) List(TripId string) ([]Frequencies, error) {

	var routes []Frequencies

	conds := db.Cond{}

	conds = db.Cond{"trip_id": TripId}

	count, err := self.C.Count(conds)

	routes = make([]Frequencies, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		route := Frequencies{}
		err := q.Next(&route)
		if err != nil {
			if err != db.ErrNoMoreRows {
				return nil, err
			}
			break
		}
		routes = append(routes, route)
	}

	return routes, nil
}

func (self *FrequenciesApi) Get(TripId string) (*Frequencies, error) {

	conds := db.Cond{"trip_id": TripId}

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	route := &Frequencies{}

	q.One(route)

	if route == nil {
		return nil, ErrNotFound
	}

	return route, nil
}
