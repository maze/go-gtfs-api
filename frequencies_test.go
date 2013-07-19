package main

import (
	"testing"
)

var api *FrequenciesApi

func TestConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(FrequenciesCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	api = &FrequenciesApi{
		DB: sess,
		C:  col,
	}

}

func TestList(t *testing.T) {
	all, err := api.List("36237")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestGetById(t *testing.T) {
	found, err := api.Get("36237")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
