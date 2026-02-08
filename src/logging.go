package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel/trace"
)

type spanContextHandler struct {
	slog.Handler
	projectID string
}

func (c spanContextHandler) Handle(ctx context.Context, r slog.Record) error {
	s := trace.SpanContextFromContext(ctx)
	if s.IsValid() {
		traceID := s.TraceID().String()
		traceID = fmt.Sprintf("projects/%s/traces/%s", c.projectID, traceID)

		r.AddAttrs(slog.String("logging.googleapis.com/trace", traceID))
		r.AddAttrs(slog.Any("logging.googleapis.com/spanId", s.SpanID()))
		r.AddAttrs(slog.Bool("logging.googleapis.com/trace_sampled", s.TraceFlags().IsSampled()))
	} else {
		fmt.Println("no span context")
	}

	return c.Handler.Handle(ctx, r)
}

func replace(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.MessageKey {
		a.Key = "message"
	}

	if a.Key == slog.LevelKey {
		a.Key = "severity"
	}

	if a.Key == slog.SourceKey {
		a.Key = "logging.googleapis.com/sourceLocation"
	}

	return a
}

func initLogging() {
	handlerOption := &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: replace,
		AddSource:   true,
	}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handlerOption)
	handler := spanContextHandler{
		Handler:   jsonHandler,
		projectID: getProjectID(),
	}
	slog.SetDefault(slog.New(handler))
}
