package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func handler(c *gin.Context) {
	requests.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
	c.String(http.StatusOK, "Hello, Grafana!")
}

func main() {
	r := gin.Default()
	r.GET("/", handler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8080")
}
