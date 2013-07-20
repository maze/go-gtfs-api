package main

import (
	"testing"
)

var api *StopApi

func TestConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(StopCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	api = &StopApi{
		DB: sess,
		C:  col,
	}

}

func TestGetById(t *testing.T) {
	found, err := api.Get("STOP_14877")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
