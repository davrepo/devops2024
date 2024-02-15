package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"net/http"

	database "minitwit.com/devops/src/database"
	flash "minitwit.com/devops/src/flash"
	model "minitwit.com/devops/src/models"
)

func GetMessages(user string, page string) []map[string]interface{} {
	var results []map[string]interface{}

	offset, messagesPerPage := LimitMessages(page)

	if user == "" {
		database.DB.Table("messages").Limit(messagesPerPage).Order("created_at desc").Offset(offset).Find(&results)
	} else {
		database.DB.Table("messages").Where("author = ?", user).Limit(messagesPerPage).Order("created_at desc").Offset(offset).Find(&results)
	}
	return results
}

func LimitMessages(page string) (int, int) {
	messagesPerPage := 50
	p, err := strconv.Atoi(page)
	if err != nil {
		panic("Failed to parse page number")
	}
	offset := (p - 1) * messagesPerPage
	return offset, messagesPerPage
}

func AddMessage(c *gin.Context) {
	user, err := c.Cookie("token")
	if err != nil || user == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	// Check if the user exists
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", user).Count(&count)
	if count == 0 {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	// Check if the message is not empty
	message := c.PostForm("message")
	if strings.TrimSpace(message) == "" {
		c.Redirect(http.StatusFound, "/user_timeline")
		return
	}

	// Create and save the message
	t := time.Now()
	database.DB.Create(&model.Message{
		Author:    user,
		Text:      message,
		CreatedAt: t,
	})
	flash.SetFlash(c, "message", "Your message was recorded")
	// Redirect to user timeline with a success message
	c.Redirect(http.StatusFound, "/user_timeline?message=success")
}

func GetFollower(follower uint, following uint) bool {
	var follows []model.Follow
	if follower == following {
		return false
	} else {
		database.DB.Find(&follows).Where("follower = ?", following).Where("following = ?", follower).First(&follows)
		return len(follows) > 0
	}
}
