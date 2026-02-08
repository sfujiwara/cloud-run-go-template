package main

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func initOpenTelemetry() (err error) {
	// Setup propagator.
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)

	// Setup tracer provider without exporter.
	tp := trace.NewTracerProvider()
	otel.SetTracerProvider(tp)

	return nil
}
