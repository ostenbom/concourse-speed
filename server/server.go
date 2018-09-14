package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleHome())
	router.HandleFunc("/api/speeddata", HandleData())

	staticFileDir := http.Dir("static")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))

	router.PathPrefix("/static/").Handler(staticFileHandler)
	return router
}

func HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeTemplate, err := ioutil.ReadFile(filepath.Join("templates", "index.html"))
		if err != nil {
			fmt.Fprintf(w, "<html><p>Error: %s</p></html>", err)
		}
		io.WriteString(w, string(homeTemplate))
	}
}

func HandleData() http.HandlerFunc {
	return func(w http.ResponseWriter, t *http.Request) {
		type Nothing struct {
			Entry string
		}
		thing := Nothing{"something"}
		thingBytes, err := json.Marshal(thing)
		if err != nil {
			fmt.Fprintf(w, "<html><p>Error: %s</p></html>", err)
		}
		w.Write(thingBytes)
	}
}
