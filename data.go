package main

import(
	"database/sql"
	"log"

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
		type TEXT NOT NULL,
		healthCheckUrl TEXT,
		PRIMARY KEY('name', 'ipAddr')
	)`
	stmt, err := self.Con.Prepare(sqlstmt)
	checkErr(err)

	stmt.Exec()
}

func (self *Database) PostService(s ServiceRegistry) error {
	sqlstmt := "INSERT INTO service_registry(name, hostname, ipAddr, status, port, type, healthCheckUrl) VALUES(?,?,?,?,?,?,?)"
	stmt, err := self.Con.Prepare(sqlstmt)
	if err != nil {
		return err
	}

	stmt.Exec(s.Name, s.Hostname, s.IpAddr, s.Status, s.Port, s.Type, s.HealthCheckUrl)
	return nil
}

func (self *Database) GetService() ([]ServiceRegistry, error){
	sqlstmt := "SELECT name, hostname, ipAddr, status, port, healthCheckUrl FROM service_registry"
	rows, err := self.Con.Query(sqlstmt)
	if err != nil {
		return []ServiceRegistry{}, err
	}

	s := ServiceRegistry{}
	list := []ServiceRegistry{}
	for rows.Next(){
		err = rows.Scan(&s.Name, &s.Hostname, &s.IpAddr, &s.Status, &s.Port, &s.HealthCheckUrl)
		if err != nil {
			return []ServiceRegistry{}, err
		}
		list = append(list, s)
	}
	rows.Close()

	return list, nil
}

func (self *Database) DeleteService(s ServiceRegistry) error {
	sqlstmt := "DELETE FROM service_registry WHERE hostname=? AND ipAddr=? AND name=?"
	stmt, err := self.Con.Prepare(sqlstmt)
	if err != nil {
		return err
	}

	stmt.Exec(s.Hostname, s.IpAddr, s.Name)
	return nil
}