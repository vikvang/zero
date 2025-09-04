package main

import (
	"log/slog"
	"net/http"
	"os"

	_ "net/http/pprof" // profiling

	_ "github.com/joho/godotenv/autoload" // automatically load .env files

	"github.com/vikvang/zero/internal/cmd"
	"github.com/vikvang/zero/internal/log"
)

func main() {
	defer log.RecoverPanic("main", func() {
		slog.Error("Application terminated due to unhandled panic")
	})

	if os.Getenv("ZERO_PROFILE") != "" {
		go func() {
			slog.Info("Serving pprof at localhost:6060")
			if httpErr := http.ListenAndServe("localhost:6060", nil); httpErr != nil {
				slog.Error("Failed to pprof listen", "error", httpErr)
			}
		}()
	}

	cmd.Execute()
}
