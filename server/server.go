package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

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
	router.HandleFunc("/api/speeddata/{period}", HandleData(db))

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
		// vars := mux.Vars(t)
		// period := vars["period"]
		weekAgo := time.Now().AddDate(0, 0, -7)
		weekAgoString := weekAgo.Format("2006-01-02 15:04:05")
		rows, err := db.Query(fmt.Sprintf(`SELECT builds.name, jobs.name, status, start_time, end_time
		FROM builds INNER JOIN jobs ON builds.job_id = jobs.id WHERE start_time > %s;`, weekAgoString))
		if err != nil {
			http.Error(w, fmt.Sprintf("Query error: %s", err), http.StatusInternalServerError)
		}

		for rows.Next() {
			var (
				build  string
				job    string
				status string
				start  time.Time
				end    time.Time
			)
			err = rows.Scan(&build, &job, &status, &start, &end)

			fmt.Println(build, job, status, start, end)
			if err != nil {
				http.Error(w, fmt.Sprintf("Query error: %s", err), http.StatusInternalServerError)
			}
		}

		type Nothing struct {
			Entry string
		}
		thing := Nothing{"something"}
		thingBytes, err := json.Marshal(thing)
		if err != nil {
			http.Error(w, fmt.Sprintf("Marshal error: %s", err), http.StatusInternalServerError)
		}
		w.Write(thingBytes)
	}
}
