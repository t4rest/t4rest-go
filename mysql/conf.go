package mysql

import "github.com/pkg/errors"

// Conf .
type Conf struct {
	ConnectionString string
	User             string
	Password         string
	MigrationDir     string
	DebugMode        bool
	MaxOpenConns     int
	MaxIdleConns     int
}

// Validate .
func (c Conf) Validate() error {

	if c.ConnectionString == "" {
		return errors.New("no mysql ConnectionString")
	}

	if c.User == "" {
		return errors.New("no mysql User")
	}

	if c.Password == "" {
		return errors.New("no mysql Password")
	}

	if c.MaxOpenConns == 0 {
		return errors.New("no mysql MaxOpenConns")
	}

	if c.MaxIdleConns == 0 {
		return errors.New("no mysql MaxIdleConns")
	}

	return nil
}
