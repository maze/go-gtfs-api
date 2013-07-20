package main

import (
	"fmt"
	"github.com/xiam/bridge"
	"menteslibres.net/gosexy/db"
)

func main() {
	var sess db.Database
	var err error

	sess, err = getDatabase()

	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	srv := bridge.New("tcp", ":11223")

	agencyContext := NewAgencyContext(sess)

	srv.AddRoute("/agency/get", (*agencyContext).Get)
	srv.AddRoute("/agency/list", (*agencyContext).List)
	srv.AddRoute("/agency/search", (*agencyContext).Search)

	err = srv.Start()

	if err != nil {
		fmt.Errorf("Error: %v\n", err)
	}
}
