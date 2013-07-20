package main

import (
	"testing"
)

var agencyApi *AgencyApi

func TestAgencyConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(AgencyCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	agencyApi = &AgencyApi{
		DB: sess,
		C:  col,
	}

}

func TestAgencyList(t *testing.T) {
	all, err := agencyApi.List()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestAgencyGet(t *testing.T) {
	found, err := agencyApi.Get("MB")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestAgencySearch(t *testing.T) {
	all, err := agencyApi.Search(`metro`)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}
