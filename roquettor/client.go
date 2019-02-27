package roquettor

import (
	"errors"
	"log"
	"strings"
)

type client struct {
	Name    string
	RClient RClient
}

var (
	hive       = client{Name: "hive", RClient: hiveClient{}}
	postgreSQL = client{Name: "postgres", RClient: defaultClient{}}
)

var clients = []client{
	hive,
	postgreSQL,
}

/* Clients impl. */
type defaultClient struct{}
type hiveClient struct{}

func (h hiveClient) query(string) int32 {
	// TODO: Implementation Hive with https://github.com/beltran/gohive
	return 1
}

func (d defaultClient) query(string) int32 {
	// TODO: Classic implementation with database/sql for native drivers
	return 0
}

// Client represents a SQL Database client
type RClient interface {
	query(string) int32 // Returns rows count
}

func NewRClient(d *Database) RClient {
	client, err := GetClientByType(d.Driver)
	if err != nil {
		log.Panic(err)
	}
	return client.RClient
}

func GetClientByType(name string) (client, error) {
	for _, c := range clients {
		if strings.EqualFold(c.Name, name) {
			return c, nil
		}
	}
	return client{}, errors.New("Type does'nt exist or not supported.")
}

// TypeExists - Returns true if type name exist
func TypeNameExists(name string) bool {
	_, err := GetClientByType(name)
	return err == nil
}
