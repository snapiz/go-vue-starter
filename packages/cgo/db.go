package cgo

import (
	"database/sql"
	"os"
	"regexp"
	"strings"
)

var dbnameRegex = regexp.MustCompile(`dbname=([a-zA-Z_0-9]+)`)
var dbSourceKey = "DATABASE_SOURCE"

// GetDBName find dbname value and return it
func GetDBName() string {
	source := os.Getenv(dbSourceKey)
	result := dbnameRegex.FindStringSubmatch(source)

	if len(result) > 0 {
		return result[1]
	}

	return ""
}

// GetDBSource get db source and return it
func GetDBSource() string {
	return os.Getenv(dbSourceKey)
}

// NewDB create new DB
func NewDB(prefix string, noDB bool) (*sql.DB, error) {
	source := os.Getenv(prefix + dbSourceKey)

	if noDB {
		result := dbnameRegex.FindStringSubmatch(source)
		source = strings.Replace(source, " "+result[0], "", 1)
	}

	return sql.Open("postgres", source)
}
