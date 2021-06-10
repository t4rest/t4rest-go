package retry

import (
	"time"
)

// Conf .
type Conf struct {
	Attempts   uint
	DelayMs    time.Duration
	MaxDelayMs time.Duration
}
