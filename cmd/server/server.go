package main

import (
	"context"
	"fmt"
	"gpg/portal/internal/localdb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func NewServer(
	ctx context.Context,
	db localdb.Db,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		ctx,
		db,
	)
	var handler http.Handler = mux
	return handler
}

func run(
	ctx context.Context,
	getenv func(string) string,
	db localdb.Db,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	initConfig()
	db.InitDb()
	srv := NewServer(ctx, db)
	httpServer := &http.Server{
		Addr:    getenv("PORT"),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := http.ListenAndServe(httpServer.Addr, httpServer.Handler); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
