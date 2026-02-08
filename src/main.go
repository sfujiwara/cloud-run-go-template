package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	slog.InfoContext(ctx, "hello")
	slog.InfoContext(ctx, "world")
	fmt.Fprintf(w, "hello")
}

func main() {
	initOpenTelemetry()
	initLogging()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/", otelhttp.NewHandler(http.HandlerFunc(handler), "handler"))

	// Start HTTP server.
	slog.Info("listening on port " + port)
	http.ListenAndServe(":"+port, nil)
}
