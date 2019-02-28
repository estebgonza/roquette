package roqclient

import "errors"

// DefaultClient - Abstract to standardize Roquettor communications with all drivers
type DefaultClient struct{}

// Connect - Default connection implementation
func (d DefaultClient) Connect(host string, port int, user string, pass string) error {
	// TODO: Classic implementation with database/sql for native drivers
	return errors.New("`Connect` is not implemented for this driver")
}

// Query - Default query implementation
func (d DefaultClient) Execute(string) (int32, error) {
	// TODO: Classic implementation with database/sql for native drivers
	return 0, errors.New("`Execute` is not implemented for this driver")
}
