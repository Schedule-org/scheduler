package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
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
// @description Saas plataforma

// @contact.name Hebert Santos
// @contact.url https://www.hebertzin.com/
// @contact.email hebertsantosdeveloper@gmail.com

// @BasePath /api/v1
func main() {

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	r := gin.Default()
	r.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", handler)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	logrus.WithFields(logrus.Fields{
		"status": "success",
	}).Info("Server is running!!")
	r.Run(":8080")

}
