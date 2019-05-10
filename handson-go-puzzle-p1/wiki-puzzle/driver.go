package driver

import (
	"database/sql"

    _ "github.com/mattn/go-sqlite3"
)

// DB ...
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// DBConn ...
var dbConn = &DB{}

// ConnectSQL ...
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	d, err := sql.Open("sqlite3", "./wiki.db")
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}
