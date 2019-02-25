package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for i := 0; i < 1000; i++ {
			opsProcessed.Inc()
			r := rand.Intn(10)
			time.Sleep(time.Duration(r) * time.Microsecond)
			opHistogram.Observe(float64(r))
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	opHistogram = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "myapp_processed_ops_hist",
		Help: "Hist for ops",
	})
)

func main() {
	var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

	recordMetrics()

	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))

}
