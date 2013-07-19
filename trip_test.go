package main

import (
	"testing"
)

var api *TripApi

func TestConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(TripCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	api = &TripApi{
		DB: sess,
		C:  col,
	}

}

func TestList(t *testing.T) {
	all, err := api.List("ROUTE_18226")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestGetById(t *testing.T) {
	found, err := api.Get("38834")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
