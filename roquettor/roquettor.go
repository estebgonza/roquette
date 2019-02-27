package roquettor

import (
	"log"

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
	if !checkInputDatabase(d) || !checkInputPlan(p) {
		log.Fatal("Please check your input files.")
		return
	}
	rclient := NewRClient(d)
	rclient.query("SHOW TABLE")
}

func checkInputDatabase(d *Database) bool {
	if d.Driver == "" {
		log.Fatal("Please specify a driver type.")
		return false
	} else if !TypeNameExists(d.Driver) {
		log.Fatal("Specified type name is not supported.")
		return false
	}
	return true
}

func checkInputPlan(p *Plan) bool {
	// TODO: Check Plan inputs
	return true
}
