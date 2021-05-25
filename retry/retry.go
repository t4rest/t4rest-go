package retry

import "github.com/avast/retry-go"

// Retry .
type Retry struct {
	cfg  Conf
	opts []retry.Option
}

//	DefaultAttempts      = uint(10)
//	DefaultDelay         = 100 * time.Millisecond
//	DefaultMaxJitter     = 100 * time.Millisecond
//	DefaultOnRetry       = func(n uint, err error) {}
//	DefaultRetryIf       = IsRecoverable
//	DefaultDelayType     = CombineDelay(BackOffDelay, RandomDelay)
//	DefaultLastErrorOnly = false
//	DefaultContext       = context.Background()

// New .
func New(cfg Conf) *Retry {
	r := &Retry{cfg: cfg}

	r.opts = make([]retry.Option, 0, 10)

	if cfg.Attempts != 0 {
		r.opts = append(r.opts, retry.Attempts(cfg.Attempts))
	}

	if cfg.Delay != 0 {
		r.opts = append(r.opts, retry.Delay(cfg.Delay))
	}

	if cfg.MaxDelay != 0 {
		r.opts = append(r.opts, retry.Delay(cfg.MaxDelay))
	}

	if cfg.MaxJitter != 0 {
		r.opts = append(r.opts, retry.Delay(cfg.MaxJitter))
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
