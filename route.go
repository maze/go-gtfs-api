package main

import (
	"menteslibres.net/gosexy/db"
)

type Route struct {
	AgencyId     string `field:"agency_id"`
	ShortName    string `field:"route_short_name"`
	LongName     string `field:"route_long_name"`
	Description  string `field:"route_desc"`
	Type         int    `field:"route_type"`
	URL          string `field:"route_url"`
	Color        string `field:"route_color"`
	TextColor    string `field:"route_text_color"`
	BikesAllowed bool   `field:"route_bikes_allowed"`
	RouteId      string `field:"route_id"`
}

const RouteCollection = `routes`

type RouteApi struct {
	C  db.Collection
	DB db.Database
}

func (self *RouteApi) List(AgencyId string) ([]Route, error) {
	var routes []Route

	conds := db.Cond{}

	conds = db.Cond{"agency_id": AgencyId}

	count, err := self.C.Count(conds)

	routes = make([]Route, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		route := Route{}
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

func (self *RouteApi) Get(RouteId string) (*Route, error) {

	conds := db.Cond{"route_id": RouteId}

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	route := &Route{}

	q.One(route)

	if route == nil {
		return nil, ErrNotFound
	}

	return route, nil
}
