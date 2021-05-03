package httpserver

import "time"

type Conf struct {
	ListenOnPort          string
	ServerReadTimeoutSec  time.Duration
	ServerWriteTimeoutSec time.Duration
	ServerIdleTimeoutSec  time.Duration
	GracefulShutdownSec   time.Duration
}
