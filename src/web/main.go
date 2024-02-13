package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/penglongli/gin-metrics/ginmetrics"
	controller "minitwit.com/devops/src/controller"
	database "minitwit.com/devops/src/database"
	model "minitwit.com/devops/src/models"
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

func main() {
	if err := godotenv.Load(".minitwit-secrets.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
		"getUserId":    GetUserID,
	})
	router.LoadHTMLGlob("src/web/templates/*.tpl")
	router.Static("/web/static", "./src/web/static")

	database.SetupDB()

	/*
		/ Shows a user's timeline or redirects to /public if no user is logged in.
		   This timeline displays the user's messages as well as all the messages of followed users.
		/public: Displays the latest messages of all users.
		/<username>: Displays a user's tweets.
		/<username>/follow: Adds the current user as a follower of the given user.
		/<username>/unfollow: Removes the current user as a follower of the given user.
		/add_message: POST endpoint to register a new message for the user.
		/login: GET and POST endpoints to log the user in.
		/register: GET and POST endpoints to register a new user.
		/logout: Logs the user out.
	*/

	router.GET("/", controller.Timeline)
	router.GET("/version", (func(c *gin.Context) {
		c.Data(200, "application/json; charset=utf-8", []byte(os.Getenv("VERSION")))
	}))
	router.GET("/public", controller.Timeline)
	router.GET("/:username", controller.UserTimeline)

	router.GET("/register", controller.Register)
	router.POST("/register", controller.SignUp)
	router.GET("/login", controller.LoginPage)
	router.POST("/login", controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/:username/follow", controller.Follow)
	router.GET("/:username/unfollow", controller.Unfollow)
	router.POST("/add_message", controller.AddMessage)

	getGinMetrics(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
