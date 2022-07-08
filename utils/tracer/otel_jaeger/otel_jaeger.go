package otel_jaeger

import (
	"context"
	"log"

	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	jaegerLib "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type Options struct {
	ServiceName string `json:"service_name" mapstructure:"service_name"`
	EndpointURL string `json:"endpoint_url" mapstructure:"endpoint_url"`
	Username    string `json:"username" mapstructure:"username"`
	Password    string `json:"password" mapstructure:"password"`
}

func NewCollector(opts Options) tracer.Provider {
	exporter, err := jaegerLib.New(jaegerLib.WithCollectorEndpoint(
		jaegerLib.WithEndpoint(opts.EndpointURL),
	))

	if err != nil {
		log.Fatal("we can't initialize jaeger tracer", err)
	}

	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			attribute.String("service.name", opts.ServiceName),
			attribute.String("library.language", "go"),
		)),
	)

	otel.SetTracerProvider(tracerProvider)
	return newProvider(tracerProvider, opts.ServiceName)
}

func newProvider(provider *tracesdk.TracerProvider, serviceName string) tracer.Provider {
	return &jaegerImpl{provider: provider, serviceName: serviceName}
}

type jaegerImpl struct {
	provider    *tracesdk.TracerProvider
	serviceName string
}

func (tr *jaegerImpl) Shutdown(ctx context.Context) error {
	return tr.provider.Shutdown(ctx)
}

func (tr *jaegerImpl) GetName() string {
	return "jaeger"
}

func (tr *jaegerImpl) GetServiceName() string {
	return tr.serviceName
}
