package main

import (
	"net/http"

	"github.com/ostenbom/concourse-speed/server"
)

func main() {
	http.HandleFunc("/", server.HandleHome("templates"))
	http.ListenAndServe(":8080", nil)
}
