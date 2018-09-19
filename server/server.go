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
			return
		}
		io.WriteString(w, string(homeTemplate))
	}
}

func HandleData(db database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, t *http.Request) {
		vars := mux.Vars(t)
		period := vars["period"]

		timeAgo, err := timeFromPeriod(period)
		if err != nil {
			http.Error(w, fmt.Sprintf("Argument error: %s", err), http.StatusInternalServerError)
		}
		timeAgoString := timeAgo.Format("2006-01-02 15:04:05")

		rows, err := db.Query(fmt.Sprintf(`SELECT builds.name, jobs.name, status, start_time, end_time
																				FROM builds INNER JOIN jobs ON builds.job_id = jobs.id
																				WHERE start_time > '%s';`, timeAgoString))
		if err != nil {
			http.Error(w, fmt.Sprintf("Query error: %s", err), http.StatusInternalServerError)
			return
		}

		type DataResponse struct {
			Build  string
			Job    string
			Status string
			Start  *time.Time
			End    *time.Time
		}

		var dataEntries []DataResponse

		for rows.Next() {
			dataEntry := DataResponse{}
			err = rows.Scan(&dataEntry.Build, &dataEntry.Job, &dataEntry.Status, &dataEntry.Start, &dataEntry.End)

			if err != nil {
				http.Error(w, fmt.Sprintf("Scan error: %s", err), http.StatusInternalServerError)
				return
			}

			dataEntries = append(dataEntries, dataEntry)
		}

		thingBytes, err := json.Marshal(dataEntries)
		if err != nil {
			http.Error(w, fmt.Sprintf("Marshal error: %s", err), http.StatusInternalServerError)
			return
		}
		w.Write(thingBytes)
	}
}

func timeFromPeriod(period string) (time.Time, error) {
	if period == "week" {
		return time.Now().AddDate(0, 0, -7), nil
	} else if period == "month" {
		return time.Now().AddDate(0, -1, 0), nil
	} else if period == "quarter" {
		return time.Now().AddDate(0, -3, 0), nil
	} else if period == "year" {
		return time.Now().AddDate(-1, 0, 0), nil
	}

	return time.Now(), fmt.Errorf("Period %s is not an accepted period", period)
}
