package main

import (
	"testing"
)

var routesApi *RouteApi

func TestRoutesConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(RouteCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	routesApi = &RouteApi{
		DB: sess,
		C:  col,
	}

}

func TestRoutesList(t *testing.T) {
	all, err := routesApi.List("MB")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestRoutesGetById(t *testing.T) {
	found, err := routesApi.Get("ROUTE_18226")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
