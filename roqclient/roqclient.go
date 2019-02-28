package roqclient

import (
	"errors"
	"strings"
)

type client struct {
	Name   string
	Client RoqClient
}

var (
	hive = client{Name: "hive", Client: &HiveClient{}}
	// postgreSQL = client{Name: "postgres", Client: DefaultClient{}}
)

var clients = []client{
	hive,
	// postgreSQL,
}

/* Clients impl. */

// RoqClient - represents a SQL Database client
type RoqClient interface {
	Connect(host string, port int, user string, pass string) error // Returns true if connection is success
	Execute(query string) (int32, error)                           // Returns rows count
}

// NewRClient - RClient factory
// Returns the right RClient interface implementation with specified Database
func NewRClient(driver string) (RoqClient, error) {
	client, err := getClientByType(driver)
	if err != nil {
		return nil, err
	}
	return client.Client, nil
}

func getClientByType(name string) (*client, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("driver name not specified")
	}
	for _, c := range clients {
		if strings.EqualFold(c.Name, name) {
			return &c, nil
		}
	}
	return &client{}, errors.New("type does'nt exist or not supported")
}
