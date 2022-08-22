package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func runAway(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusOK,
	})
}

func runAway2(c *gin.Context) {
	c.JSON(http.StatusBadGateway, gin.H{
		"message": http.StatusBadGateway,
	})
}

func main() {
	r := gin.Default()
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	m.SetMetricPath("/metrics")
	m.SetSlowTime(10) // +optional set slow time, default 5s
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)

	r.GET("/endpoint1", runAway)
	r.GET("/endpoint2", runAway)
	r.POST("/endpoint3", runAway)
	r.OPTIONS("/endpoint4", runAway)
	r.Run()
}
