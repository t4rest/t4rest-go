package main

import (
	"github.com/pkg/errors"

	"github.com/t4rest/t4rest-go/logger"
	"github.com/t4rest/t4rest-go/module"
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

	module.Run(log)
}
