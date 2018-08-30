package main

import (
	"net/http"

	"github.com/ostenbom/concourse-speed/server"
)

func main() {
	http.HandleFunc("/", server.HandleHome())
	http.ListenAndServe(":8080", nil)
}
