package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ostenbom/concourse-speed/server"
)

var _ = Describe("Server", func() {
	var (
		server *httptest.Server
	)
	BeforeEach(func() {
		server = httptest.NewServer(HandleHome("../templates/"))
	})

	Context("when visiting home page", func() {
		It("returns a response", func() {
			body := getResponseBody(server)
			Expect(len(body)).NotTo(BeZero())
		})

		It("returns html, body, doctype", func() {
			body := getResponseBody(server)
			Expect(body).To(ContainSubstring("<html>"))
			Expect(body).To(ContainSubstring("</html>"))
			Expect(body).To(ContainSubstring("<body>"))
			Expect(body).To(ContainSubstring("</body>"))
			Expect(body).To(ContainSubstring("<!doctype html>"))
		})

		It("has head and title Concourse Speed", func() {
			body := getResponseBody(server)
			Expect(body).To(ContainSubstring("<head>"))
			Expect(body).To(ContainSubstring("</head>"))
			Expect(body).To(ContainSubstring("<title>Concourse Speed</title>"))
		})
	})

	AfterEach(func() {
		server.Close()
	})

})

func getResponseBody(server *httptest.Server) string {
	response, err := http.Get(server.URL)
	Expect(err).NotTo(HaveOccurred())

	body, err := ioutil.ReadAll(response.Body)
	Expect(err).NotTo(HaveOccurred())

	return string(body)
}
