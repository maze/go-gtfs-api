package main

import (
	"testing"
)

var shapesApi *ShapeApi

func TestShapesConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(ShapeCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	shapesApi = &ShapeApi{
		DB: sess,
		C:  col,
	}

}

func TestShapesList(t *testing.T) {
	all, err := shapesApi.List("36477")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}
