package server

import (
	"io"
	"net/http"
)

func HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><title>Concourse Speed</title><body><h1>Hello!</h1></body></html>")
	}
}
