package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(requests)
}

func handler(w http.ResponseWriter, r *http.Request) {
	requests.WithLabelValues(r.Method, r.URL.Path).Inc()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Grafana!"))
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
