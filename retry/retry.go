package retry

import "github.com/avast/retry-go"

// Retry .
type Retry struct {
	cfg  Conf
	opts []retry.Option
}

// New .
func New(cfg Conf) *Retry {
	r := &Retry{cfg: cfg}

	r.opts = make([]retry.Option, 0, 10)

	if cfg.Attempts != 0 {
		r.opts = append(r.opts, retry.Attempts(cfg.Attempts))
	}

	if cfg.DelayMs != 0 {
		r.opts = append(r.opts, retry.Delay(cfg.DelayMs))
	}

	if cfg.MaxDelayMs != 0 {
		r.opts = append(r.opts, retry.Delay(cfg.MaxDelayMs))
	}

	return r
}

// Do .
func (r *Retry) Do(action retry.RetryableFunc, ops ...retry.Option) error {

	if len(ops) > 0 {
		r.opts = append(r.opts, ops...)
	}

	err := retry.Do(action, r.opts...)

	return err
}
