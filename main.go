package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ostenbom/concourse-speed/server"
	"github.com/ostenbom/concourse-speed/server/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pivotal"
	password = "***"
	dbname   = "atc"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)
	database, err := database.New(psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not set up router: %s", err)
	}

	router, err := server.NewRouter(database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not set up router: %s", err)
	}
	http.ListenAndServe(":8080", router)
}
