package main

import (
	"menteslibres.net/gosexy/db"
)

type Trip struct {
	RouteId            string `field:"route_id"`
	ServiceId          string `field:"service_id"`
	TripShortName      string `field:"trip_short_name"`
	TripHeadsign       string `field:"trip_headsign"`
	RouteShortName     string `field:"route_short_name"`
	DirectionId        bool   `field:"direction_id"`
	BlockId            string `field:"block_id"`
	ShapeId            string `field:"shape_id"`
	WeelChairAccesible bool   `field:"weelchair_accesible"`
	TripBikesAllowed   bool   `field:"trip_bikes_allowed"`
	TripId             int    `field:"trip_id"`
}

const TripCollection = `trips`

type TripApi struct {
	C  db.Collection
	DB db.Database
}

func (self *TripApi) List(RouteId string) ([]Trip, error) {
	var routes []Trip

	conds := db.Cond{}

	conds = db.Cond{"route_id": RouteId}

	count, err := self.C.Count(conds)

	routes = make([]Trip, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		route := Trip{}
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

func (self *TripApi) Get(RouteId string) (*Trip, error) {

	conds := db.Cond{"route_id": RouteId}

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	route := &Trip{}

	q.One(route)

	if route == nil {
		return nil, ErrNotFound
	}

	return route, nil
}
