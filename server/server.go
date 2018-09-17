package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/ostenbom/concourse-speed/server/database"
)

func NewRouter(db database.Database) (*mux.Router, error) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleHome())

	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not query database: %s", err)
	}
	router.HandleFunc("/api/speeddata", HandleData(db))

	staticFileDir := http.Dir("static")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))

	router.PathPrefix("/static/").Handler(staticFileHandler)
	return router, nil
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

func HandleData(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, t *http.Request) {
		db.Query("haha")
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
