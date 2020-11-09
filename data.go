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
}