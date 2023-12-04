package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "handling incoming request")
		defer func() {
		slog.InfoContext(r.Context(), "handled incoming request")
		}()

		fmt.Fprintf(w, "Hello, World!")
	})

	slog.InfoContext(context.TODO(), "starting server")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		slog.ErrorContext(context.TODO(), "failed to start server", slog.String("error", err.Error()))
		panic(err)
	}
}
