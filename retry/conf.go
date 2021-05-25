package retry

import (
	"time"
)

// Conf .
type Conf struct {
	Attempts  uint
	Delay     time.Duration
	MaxDelay  time.Duration
	MaxJitter time.Duration
}
