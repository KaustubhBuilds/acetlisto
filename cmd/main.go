package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KaustubhBuilds/acetlisto/handlers"
	"github.com/KaustubhBuilds/acetlisto/stores"
	gorillahandlers "github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

func main() {
	// -----------------------------
	// 1. OpenTelemetry INIT
	// -----------------------------
	ctx := context.Background()

	var shutdown func(context.Context) error

	if os.Getenv("ENABLE_CLOUD_TRACE") == "true" {
		var err error
		shutdown, err = initTracer(ctx)
		if err != nil {
			log.Fatalf("failed to initialize cloud tracing: %v", err)
		}

		defer func() {
			if err := shutdown(ctx); err != nil {
				log.Printf("failed to shut down tracer: %v", err)
			}
		}()

		log.Println("Google Cloud Trace enabled")
	} else {
		log.Println("Google Cloud Trace disabled")
	}

	// -----------------------------
	// 2. Config & Store
	// -----------------------------
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("READ_TIMEOUT", "30")

	store := stores.NewMemoryItemStore()

	// -----------------------------
	// 3. Router Setup
	// -----------------------------
	// handlers.New() wraps the router with otelmux middleware
	r := handlers.New(store)

	// -----------------------------
	// 4. Server Start
	// -----------------------------
	port := viper.GetString("PORT")
	log.Printf("Server starting on :%s", port)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", port),
		Handler:     gorillahandlers.LoggingHandler(os.Stdout, r),
		ReadTimeout: viper.GetDuration("READ_TIMEOUT") * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
