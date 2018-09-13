package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ostenbom/concourse-speed/server"
)

var _ = Describe("Server", func() {
	var (
		server *httptest.Server
	)
	BeforeEach(func() {
		_, filename, _, _ := runtime.Caller(0)
		rootPath := path.Join(path.Dir(filename), "..")
		os.Chdir(rootPath)
		server = httptest.NewServer(NewRouter())
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

	AfterEach(func() {
		server.Close()
	})

})

func getResponseBody(server *httptest.Server, path string) string {
	response, err := http.Get(server.URL + path)
	Expect(err).NotTo(HaveOccurred())

	Expect(response.StatusCode).To(Equal(http.StatusOK))

	body, err := ioutil.ReadAll(response.Body)
	Expect(err).NotTo(HaveOccurred())

	return string(body)
}
