package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	svr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := svr.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	select {
	case <-ctx.Done():
		stop()
		if err := svr.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
		log.Fatal("terminating: by signal ")
	}
	os.Exit(0)
}
