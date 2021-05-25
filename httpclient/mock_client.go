package httpclient

import (
	"net/http"
)

// RoundTripFunc is the mock roundTrip
type RoundTripFunc func(r *http.Request) (*http.Response, error)

func (s RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}
