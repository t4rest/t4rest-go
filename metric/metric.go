package metric

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/resource"
)

// Metric .
type Metric struct {
	exporter *prometheus.Exporter
}

// Conf .
type Conf struct {
	AppID string
}

// New .
func New(cfg Conf) (Metric, error) {
	m := Metric{}

	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(attribute.String("service", cfg.AppID)),
	)
	if err != nil {
		return m, errors.Wrap(err, "failed to create prometheus resource")
	}

	m.exporter, err = prometheus.InstallNewPipeline(
		prometheus.Config{},
		controller.WithResource(res),
	)
	if err != nil {
		return m, errors.Wrap(err, "failed to initialize prometheus exporter")
	}

	return m, nil
}

// HTTPHandlerFunc .
func (m Metric) HTTPHandlerFunc() http.HandlerFunc {
	return m.exporter.ServeHTTP
}

// HTTPRouterHandler .
func (m Metric) HTTPRouterHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		m.HTTPHandlerFunc().ServeHTTP(w, r)
	}
}
