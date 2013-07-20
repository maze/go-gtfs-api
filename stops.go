package main

import (
	"menteslibres.net/gosexy/db"
)

type Stop struct {
	StopId             string  `stop_id`
	Code               string  `stop_code`
	Name               string  `stop_name`
	Desc               string  `stop_desc`
	Lat                float64 `stop_lat`
	Lng                float64 `stop_lon`
	ZoneId             string  `zone_id`
	URL                string  `stop_url`
	LocationType       bool    `location_type`
	ParentStation      string  `parent_station`
	WheelchairBoarding string  `wheelchair_boarding`
	StopDirection      string  `stop_direction`
}

const StopCollection = `stops`

type StopApi struct {
	C  db.Collection
	DB db.Database
}

func (self *StopApi) Get(StopId string) (*Stop, error) {

	conds := db.Cond{"stop_id": StopId}

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	stop := &Stop{}

	q.One(stop)

	if stop == nil {
		return nil, ErrNotFound
	}

	return stop, nil
}
