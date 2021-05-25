package main

import (
	"context"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/errors"

	"github.com/t4rest/t4rest-go/httpclient"
	"github.com/t4rest/t4rest-go/logger"
	"github.com/t4rest/t4rest-go/module"
	ret "github.com/t4rest/t4rest-go/retry"
	"github.com/t4rest/t4rest-go/trace"
)

func main() {

	log := logger.New(logger.Conf{AppID: "AppID", LogLevel: "debug"})
	defer log.Flush()

	flush, err := trace.New(trace.Conf{AppID: "AppID", CollectorEndpoint: "http://localhost:14268/api/traces"})
	if err != nil {
		log.Fatal(errors.Wrap(err, "trace.New"))
	}
	defer flush()

	ctx := context.Background()

	httpClient := httpclient.New(httpclient.Conf{HTTPTimeout: 10 * time.Second})

	// retry example
	err500 := errors.New("Server error")

	r := ret.New(ret.Conf{Attempts: 3, Delay: 10 * time.Second, MaxDelay: 30 * time.Second})
	err = r.Do(
		func() error {

			req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := httpClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close() //nolint: errcheck

			if resp.StatusCode >= http.StatusInternalServerError {
				return errors.Wrap(err500, "server is down")
			}

			return nil
		},
		retry.OnRetry(func(i uint, err error) {
			log.With("retry_attempt", i, "error", err).Info("onRetry")
		}),
		retry.RetryIf(func(err error) bool {
			return errors.Is(err, err500)
		}),
		retry.Context(ctx),
	)
	if err != nil {
		log.With("error", err).Error("Get(google.com)")
	}

	module.Run(log)
}
