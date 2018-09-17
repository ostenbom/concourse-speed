package database_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ostenbom/concourse-speed/server/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pivotal"
	password = ""
	dbname   = "atc"
)

var _ = Describe("Database", func() {
	var database *PostgresDatabase
	BeforeEach(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable", host, port, user, dbname)

		var err error
		database, err = New(psqlInfo)
		Expect(err).NotTo(HaveOccurred())
	})

	It("can ping", func() {
		Expect(database.Ping()).To(Succeed())
	})

	Context("when querying", func() {
		It("gives me rows", func() {
			rows, err := database.Query("SELECT * FROM builds LIMIT 2;")
			Expect(err).NotTo(HaveOccurred())
			Expect(rows.Next()).To(BeTrue())
			Expect(rows.Next()).To(BeTrue())
		})
	})

})
