package main

import (
	"menteslibres.net/gosexy/db"
)

type StopTime struct {
	TripId            string  `field:"trip_id"`
	StopSequence      int     `field:"stop_sequence"`
	StopId            string  `field:"stop_id"`
	ArrivalTime       string  `field:"arrival_time"`
	DepartureTime     string  `field:"departure_time"`
	StopHeadsign      string  `field:"stop_headsign"`
	RouteShortName    string  `field:"route_short_name"`
	PickupType        int     `field:"pickup_type"`
	DropOffType       int     `field:"drop_off_type"`
	ShapeDistTraveled float64 `field:"shape_dist_traveled"`
}

const StopTimeCollection = `stop_times`

type StopTimeApi struct {
	C  db.Collection
	DB db.Database
}

func (self *StopTimeApi) List(TripId string) ([]StopTime, error) {
	var stop_times []StopTime

	conds := db.Cond{}

	conds = db.Cond{"trip_id": TripId}

	count, err := self.C.Count(conds)

	stop_times = make([]StopTime, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		stop_time := StopTime{}
		err := q.Next(&stop_time)
		if err != nil {
			if err != db.ErrNoMoreRows {
				return nil, err
			}
			break
		}
		stop_times = append(stop_times, stop_time)
	}

	return stop_times, nil
}
