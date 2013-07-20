package main

import (
	"testing"
)

var calendarApi *CalendarApi

func TestCalendarConnect(t *testing.T) {

	sess, err := getDatabase()

	if err != nil {
		t.Fatalf(err.Error())
	}

	col, err := sess.Collection(CalendarCollection)
	if err != nil {
		t.Fatalf(err.Error())
	}

	calendarApi = &CalendarApi{
		DB: sess,
		C:  col,
	}

}

func TestCalendarList(t *testing.T) {
	all, err := calendarApi.List("36238")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(all) == 0 {
		t.Fatalf(`Expecting some rows.`)
	}
}

func TestCalendarGetById(t *testing.T) {
	found, err := calendarApi.Get("36238")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if found == nil {
		t.Fatalf(`Expecting some rows.`)
	}
}
