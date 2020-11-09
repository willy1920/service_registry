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

	sqlstmt := `CREATE TABLE IF NOT EXISTS service_registry(
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

func (self *Database) PostService(s ServiceRegistry) error {
	sqlstmt := "INSERT INTO service_registry(name, hostname, ipAddr, status, port, healthCheckUrl) VALUES(?,?,?,?,?,?)"
	stmt, err := self.Con.Prepare(sqlstmt)
	if err != nil {
		return err
	}

	stmt.Exec(s.Name, s.Hostname, s.IpAddr, s.Status, s.Port, s.HealthCheckUrl)
	return nil
}