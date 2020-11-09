package main

import(
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct{
	Con *sql.DB
}

func (self *Database) InitData() {
	var err error
	self.Con, err = sql.Open("sqlite3", "data.sqlite3")
	checkErr(err)

	sqlstmt := `CREATE TABLE service_registry(
		name TEXT NOT NULL,
		hostname TEXT NOT NULL,
		ipAddr TEXT NOT NULL,
		status TEXT NOT NULL,
		port INT NOT NULL,
		healthCheckUrl TEXT NOT NULL,
		PRIMARY KEY('name', 'ipAddr')
	)`
	stmt, err := self.Con.Prepare(sqlstmt)
	checkErr(err)

	stmt.Exec()
}