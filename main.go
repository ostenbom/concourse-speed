package main

import (
	"net/http"

	"github.com/ostenbom/concourse-speed/server"
)

func main() {
	router := server.NewRouter()
	http.ListenAndServe(":8080", router)
}
