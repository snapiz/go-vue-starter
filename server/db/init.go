package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	// DB from database/sql
	DB *sql.DB
	// Fixtures fake data
	Fixtures *testfixtures.Context
)

func init() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "test" {
		godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/snapiz/go-vue-starter/.env." + appEnv))
	}

	if appEnv != "production" {
		godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/snapiz/go-vue-starter/.env.local"))
		godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/snapiz/go-vue-starter/.env"))
	}

	u, err := url.Parse(os.Getenv("DATABASE_URL"))
	s := ""
	m, _ := url.ParseQuery(u.RawQuery)

	if val, ok := m["sslmode"]; ok {
		s = " sslmode=" + val[0]
	}

	p, _ := u.User.Password()
	dbname := u.Path[1:]

	dburl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s%s", u.Hostname(), u.Port(), u.User.Username(), p, dbname, s)
	DB, err = sql.Open("postgres", dburl)

	if err != nil {
		panic(err)
	}

	boil.SetDB(DB)

	if appEnv == "test" {
		Fixtures, err = testfixtures.NewFolder(DB, &testfixtures.PostgreSQL{
			SkipResetSequences: true,
		}, os.Getenv("GOPATH")+"/src/github.com/snapiz/go-vue-starter/server/db/fixtures")

		if err != nil {
			log.Fatal(err)
		}
	}
}
