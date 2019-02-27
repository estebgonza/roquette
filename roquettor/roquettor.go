package roquettor

import (
	"fmt"
	"log"

	"database/sql"
	// Posgresql driver
	_ "github.com/lib/pq"
)

// Database - Database settings connection
type Database struct {
	Driver        string `json:"driver"`
	URIConnection string `json:"uri-connection"`
}

// Plan - Execution plan for roquettor
type Plan struct {
	Name            string `json:"name"`
	ConcurrentLevel int    `json:"concurrent-level"`
	Queries         []struct {
		SQL    string `json:"sql"`
		Repeat int    `json:"repeat"`
	} `json:"queries"`
}

// Row - Abstract row
type Row struct {
}

// Execute - test
func Execute(d *Database, p *Plan) {
	connStr := d.URIConnection
	fmt.Println(connStr)

	db, err := sql.Open(d.Driver, d.URIConnection)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SHOW TABLES")
	var c1 string
	switch err := rows.Scan(&c1); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(c1)
	default:
		panic(err)
	}
}
