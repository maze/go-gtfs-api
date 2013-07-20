package main

import (
	"testing"
)

var tripApi *TripApi

func TestTripConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(TripCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	tripApi = &TripApi{
		DB: sess,
		C:  col,
	}

}

func TestTripList(t *testing.T) {
	all, err := tripApi.List("ROUTE_18226")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestTripGetById(t *testing.T) {
	found, err := tripApi.Get("38834")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
