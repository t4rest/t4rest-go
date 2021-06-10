package httpclient

import (
	"errors"
	"time"
)

// Conf http client configuration
type Conf struct {
	HTTPTimeout time.Duration
}

// Validate .
func (c Conf) Validate() error {

	if c.HTTPTimeout == 0 {
		return errors.New("no HTTPTimeout")
	}

	return nil
}
