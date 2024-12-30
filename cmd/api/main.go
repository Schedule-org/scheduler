package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title Scheduler app
// @version 1.0
// @description Saas plataform

// @contact.name Hebert santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {
	r := gin.Default()
	r.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", handler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8080")

}
