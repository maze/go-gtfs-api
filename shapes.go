package main

import (
	"menteslibres.net/gosexy/db"
)

type Shape struct {
	ShapeId      string  `field:"shape_id"`
	PtSequence   int     `field:"shape_pt_sequence"`
	DistTraveled float64 `field:"shape_dist_traleved"`
	PtLat        float64 `field:"shape_pt_lat"`
	PtLon        float64 `field:"shape_pt_lon"`
}

const ShapeCollection = `shapes`

type ShapeApi struct {
	C  db.Collection
	DB db.Database
}

func (self *ShapeApi) List(ShapeId string) ([]Shape, error) {
	return self.query(db.Cond{"shape_id": ShapeId})
}

func (self *ShapeApi) query(conds db.Cond) ([]Shape, error) {
	var shapes []Shape

	count, err := self.C.Count(conds)

	shapes = make([]Shape, 0, count)

	q, err := self.C.Query(conds)

	if err != nil {
		return nil, err
	}

	for {
		shape := Shape{}
		err := q.Next(&shape)
		if err != nil {
			if err != db.ErrNoMoreRows {
				return nil, err
			}
			break
		}
		shapes = append(shapes, shape)
	}

	return shapes, nil
}
