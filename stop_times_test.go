package main

import (
	"testing"
)

var stopTimesApi *StopTimeApi

func TestStopTimesConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(StopTimeCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	stopTimesApi = &StopTimeApi{
		DB: sess,
		C:  col,
	}

}

func TestStopTimesList(t *testing.T) {
	all, err := stopTimesApi.List("38834")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}
