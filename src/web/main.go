package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	model "github.com/ahmedaabouzied/minitwit/src/models"
)

func getGinMetrics(router *gin.Engine) {
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/ginmetrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(router)
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	return fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", day, month, year, hour, minute, second)
}

func GetUserID(username string) uint {
	var user model.User
	database.DB.Where("username = ?", username).First(&user) // SELECT * FROM USERS WHERE USERNAME = "?"
	return user.ID
}
