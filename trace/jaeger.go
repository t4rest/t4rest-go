package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Conf .
type Conf struct {
	AppID             string
	CollectorEndpoint string
}

// New creates a new trace provider instance and registers it as global trace provider.
func New(cfg Conf) (func(), error) {
	return jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(cfg.CollectorEndpoint),
		jaeger.WithSDKOptions(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(resource.NewWithAttributes(
				semconv.ServiceNameKey.String(cfg.AppID),
			)),
		),
	)
}

////////////////////////////////////// examples ///////////////////////////////////////////
// nolint
func foo() {
	tr := otel.Tracer("component-foo")
	ctx, span := tr.Start(context.Background(), "foo")
	defer span.End()

	bar(ctx)
}

// nolint
var testKey attribute.Key = "testset"

// nolint
func bar(ctx context.Context) {
	tr := otel.Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	span.SetAttributes(testKey.String("value"), testKey.Int64(64))
	defer span.End()

	// Do bar...
}
