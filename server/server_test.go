package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ostenbom/concourse-speed/server"
	"github.com/ostenbom/concourse-speed/server/database/databasefakes"
)

var _ = Describe("Server", func() {
	var (
		server   *httptest.Server
		database *databasefakes.FakeDatabase
		rows     *databasefakes.FakeRows
	)
	BeforeEach(func() {
		_, filename, _, _ := runtime.Caller(0)
		rootPath := path.Join(path.Dir(filename), "..")
		os.Chdir(rootPath)

		database = new(databasefakes.FakeDatabase)
		router, err := NewRouter(database)
		Expect(err).NotTo(HaveOccurred())
		Expect(database.PingCallCount()).To(Equal(1))
		server = httptest.NewServer(router)
	})

	Context("when visiting home page", func() {
		It("returns a response", func() {
			body := getResponseBody(server, "/")
			Expect(len(body)).NotTo(BeZero())
		})

		It("returns html, body, doctype", func() {
			body := getResponseBody(server, "/")
			Expect(body).To(ContainSubstring("<html>"))
			Expect(body).To(ContainSubstring("</html>"))
			Expect(body).To(ContainSubstring("<body>"))
			Expect(body).To(ContainSubstring("</body>"))
			Expect(body).To(ContainSubstring("<!doctype html>"))
		})

		It("has head and title Concourse Speed", func() {
			body := getResponseBody(server, "/")
			Expect(body).To(ContainSubstring("<head>"))
			Expect(body).To(ContainSubstring("</head>"))
			Expect(body).To(ContainSubstring("<title>Concourse Speed</title>"))
		})

		It("loads speedmap.js and speedmap style", func() {
			body := getResponseBody(server, "/")
			Expect(body).To(ContainSubstring("<script src=\"/static/speedmap.js\">"))
			Expect(body).To(ContainSubstring("<link rel=\"stylesheet\" href=\"/static/speedmap.css\">"))
		})

		It("loads d3", func() {
			body := getResponseBody(server, "/")
			Expect(body).To(ContainSubstring("<script src=\"/static/d3.min.js\">"))
		})

		It("renders the chart", func() {
			body := getResponseBody(server, "/")
			Expect(body).To(ContainSubstring("<div id=\"speedmap\"></div>"))
			Expect(body).To(ContainSubstring("speedMap().render()"))
		})
	})

	Context("when serving static files", func() {
		It("serves speedmap.js", func() {
			getResponseBody(server, "/static/speedmap.js")
		})

		It("serves speedmap.css", func() {
			getResponseBody(server, "/static/speedmap.css")
		})

		It("serves d3", func() {
			getResponseBody(server, "/static/d3.min.js")
		})
	})

	Context("when getting data from /api/speeddata", func() {
		BeforeEach(func() {
			rows = new(databasefakes.FakeRows)
			rows.NextReturns(false)
			database.QueryReturns(rows, nil)
		})

		It("returns a successful response", func() {
			getResponseBody(server, "/api/speeddata/week")
		})

		It("returns a JSON response", func() {
			type JSONObj map[string]interface{}
			var responseJSON JSONObj
			response := getResponseBytes(server, "/api/speeddata/week")
			Expect(json.Unmarshal(response, &responseJSON)).To(Succeed())
		})

		It("queries the database with the right query", func() {
			getResponseBody(server, "/api/speeddata/week")
			Expect(database.QueryCallCount()).To(Equal(1))
			calledQuery, _ := database.QueryArgsForCall(0)
			Expect(calledQuery).To(ContainSubstring("SELECT builds.name, jobs.name, status, start_time, end_time"))
			Expect(calledQuery).To(ContainSubstring("FROM builds INNER JOIN jobs ON builds.job_id = jobs.id"))
		})

		Context("when querying with different periods", func() {
			It("queries a week on week", func() {
				getResponseBody(server, "/api/speeddata/week")
				weekAgo := time.Now().AddDate(0, 0, -7)
				weekAgoString := weekAgo.Format("2006-01-02 15:04:05")

				calledQuery, _ := database.QueryArgsForCall(0)
				Expect(calledQuery).To(ContainSubstring(weekAgoString))
			})
		})

	})

	AfterEach(func() {
		server.Close()
	})

})

func getResponseBytes(server *httptest.Server, path string) []byte {
	response, err := http.Get(server.URL + path)
	Expect(err).NotTo(HaveOccurred())

	Expect(response.StatusCode).To(Equal(http.StatusOK))

	body, err := ioutil.ReadAll(response.Body)
	Expect(err).NotTo(HaveOccurred())

	return body
}

func getResponseBody(server *httptest.Server, path string) string {
	body := getResponseBytes(server, path)

	return string(body)
}
