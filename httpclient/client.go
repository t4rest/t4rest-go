package httpclient

import (
	"net/http"
	"time"
)

// New http client
func New(cfg Conf) *http.Client {

	cfgDef := Conf{
		HTTPTimeout: 30 * time.Second,
	}

	if cfg.HTTPTimeout != 0 {
		cfgDef.HTTPTimeout = cfg.HTTPTimeout
	}

	return &http.Client{Timeout: cfgDef.HTTPTimeout * time.Second}
}
