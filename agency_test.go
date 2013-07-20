package main

import (
	"testing"
)

var api *AgencyApi

func TestConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(AgencyCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	api = &AgencyApi{
		DB: sess,
		C:  col,
	}

}

func TestList(t *testing.T) {
	all, err := api.List()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestSearch(t *testing.T) {
	all, err := api.Search(`metro`)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}
