package roqclient

import (
	"context"

	"github.com/beltran/gohive"
)

// HiveClient - Is a RoqClient implementation with Gohive
// Gohive is not compatible with database/sql module
// https://github.com/beltran/gohive
type HiveClient struct {
	conn *gohive.Connection
}

// Execute -
func (h *HiveClient) Execute(query string) (int32, error) {
	ctx := context.Background()
	cursor := h.conn.Cursor()
	cursor.Execute(ctx, query, false)
	return 0, cursor.Err
}

// Connect - Only support SASL connection
// TODO: Support all connections systems.
func (h *HiveClient) Connect(host string, port int, user string, pass string, db string) error {
	configuration := gohive.NewConnectConfiguration()
	configuration.Username = user
	configuration.Password = pass
	conn, err := gohive.Connect(host, port, "NONE", configuration)
	h.conn = conn
	return err
}
