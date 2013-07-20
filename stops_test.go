package main

import (
	"testing"
)

var stopsApi *StopApi

func TestStopsConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(StopCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	stopsApi = &StopApi{
		DB: sess,
		C:  col,
	}

}

func TestStopsGetById(t *testing.T) {
	found, err := stopsApi.Get("STOP_14877")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
