package main

import (
	"testing"
)

var frequenciesApi *FrequenciesApi

func TestFrequenciesConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(FrequenciesCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	frequenciesApi = &FrequenciesApi{
		DB: sess,
		C:  col,
	}

}

func TestFrequenciesList(t *testing.T) {
	all, err := frequenciesApi.List("36237")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestFrequenciesGetById(t *testing.T) {
	found, err := frequenciesApi.Get("36237")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
