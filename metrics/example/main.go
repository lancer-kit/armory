package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/lancer-kit/armory/metrics"
)

var Metrics = new(metrics.SafeMetrics)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	Metrics = Metrics.New(ctx)
	Metrics.PrettyPrint = true

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		Metrics.Collect()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		runHTTPServer(ctx)
	}()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	<-gracefulStop

	cancel()

	wg.Wait()
}

func runHTTPServer(ctx context.Context) {
	mux := http.DefaultServeMux
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		body, _ := Metrics.MarshalJSON()
		_, _ = w.Write(body)
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		Metrics.Add(metrics.NewMKey("login", r.Method))
	})
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		Metrics.Add(metrics.NewMKey("signup", r.Method))
	})
	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		key := metrics.MKey("info")
		Metrics.Add(key)
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go server.ListenAndServe()
	<-ctx.Done()
	ctxT, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	server.Shutdown(ctxT)
}
