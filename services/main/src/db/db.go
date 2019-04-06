package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/snapiz/go-vue-starter/packages/cgo"
	"github.com/volatiletech/sqlboiler/boil"
	"gopkg.in/testfixtures.v2"
)

var (
	// DB for global app
	DB *sql.DB

	// Fixtures fake data
	Fixtures *testfixtures.Context
)

func init() {
	var err error
	DB, err = cgo.NewDB("", false)

	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(DB)

	if os.Getenv("GO_ENV") == "dev" {
		boil.DebugMode = true
	}

	if os.Getenv("GO_ENV") == "test" {
		Fixtures, err = testfixtures.NewFiles(DB, &testfixtures.PostgreSQL{
			SkipResetSequences: true,
		}, "../../fixtures/users.yml")

		if err != nil {
			log.Fatal(err)
		}
	}
}
