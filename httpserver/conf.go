package httpserver

import (
	"time"

	"github.com/pkg/errors"
)

type Conf struct {
	ListenOnPort          string
	ServerReadTimeoutSec  time.Duration
	ServerWriteTimeoutSec time.Duration
	ServerIdleTimeoutSec  time.Duration
	GracefulShutdownSec   time.Duration
}

// Validate .
func (c Conf) Validate() error {

	if c.ListenOnPort == "" {
		return errors.New("no ListenOnPort")
	}

	if c.ServerReadTimeoutSec == 0 {
		return errors.New("no ServerReadTimeoutSec")
	}

	if c.ServerWriteTimeoutSec == 0 {
		return errors.New("no ServerWriteTimeoutSec")
	}

	if c.ServerIdleTimeoutSec == 0 {
		return errors.New("no ServerIdleTimeoutSec")
	}

	if c.GracefulShutdownSec == 0 {
		return errors.New("no GracefulShutdownSec")
	}

	return nil
}
