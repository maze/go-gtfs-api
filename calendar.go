package main

import (
	"menteslibres.net/gosexy/db"
)

type Calendar struct {
	ServiceId string `field:"service_id"`

	Monday    bool `field:"monday"`
	Tuesday   bool `field:"tuesday"`
	Wednesday bool `field:"wednesday"`
	Thursday  bool `field:"thursday"`
	Friday    bool `field:"friday"`
	Saturday  bool `field:"saturday"`
	Sunday    bool `field:"sunday"`

	StartDate string `field:"start_date"`
	EndDate   string `field:"end_date"`
}

const CalendarCollection = `calendar`

type CalendarApi struct {
	C  db.Collection
	DB db.Database
}

func (self *CalendarApi) List(ServiceId string) ([]Calendar, error) {
	var routes []Calendar

	conds := db.Cond{}

	conds = db.Cond{"service_id": ServiceId}

	count, err := self.C.Count(conds)

	routes = make([]Calendar, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		route := Calendar{}
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

func (self *CalendarApi) Get(ServiceId string) (*Calendar, error) {

	conds := db.Cond{"service_id": ServiceId}

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	route := &Calendar{}

	q.One(route)

	if route == nil {
		return nil, ErrNotFound
	}

	return route, nil
}
