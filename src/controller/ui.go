package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	flash "minitwit.com/devops/src/flash"
)

func Timeline(c *gin.Context) {
	user, _ := c.Cookie("token")

	page := c.DefaultQuery("page", "0")

	data := make(map[string]interface{})
	data["flashes"] = flash.GetFlash(c, "message")

	if user == "" {
		c.HTML(http.StatusOK, "timeline.tpl", gin.H{
			"title": "Timeline",
			// "endpoint": "/public_timeline",
			"flashes": data,
			"messages": GetMessages("", page, c),
		})
	} else {
		c.HTML(http.StatusOK, "timeline.tpl", gin.H{
			"title":         "Timeline",
			"user":          user,
			"user_timeline": false,
			"flashes": 		 data,
			"messages":      GetMessages("", page, c),
		})
	}
}

func UserTimeline(c *gin.Context) {
	user_query := c.Request.URL.Query().Get("username")
	data := make(map[string]interface{})
	data["flashes"] = flash.GetFlash(c, "message")
	page := c.DefaultQuery("page", "0")

	if user_query != "" {
		user, err := c.Cookie("token")
		if user != "" || err != nil {
			followed := GetFollower(GetUser(user_query).ID, GetUser(user).ID)
			var user_page = false
			if user == user_query {
				user_page = true
			}
			c.HTML(http.StatusOK, "timeline.tpl", gin.H{
				"title":         user_query + "'s Timeline",
				"user_timeline": true,
				"private":       true,
				"user":          user_query,
				"followed":      followed,
				"user_page":     user_page,
				"flashes": 		 data,
				"messages":      GetMessages(user_query, page, c),
			})
		} else {
			c.HTML(http.StatusOK, "timeline.tpl", gin.H{
				"title":         user_query + "'s Timeline",
				"user_timeline": true,
				"private":       true,
				"flashes": 		 data,
				"messages":      GetMessages(user_query, page, c),
			})
		}
	} else {
		user, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/public_timeline")
		}
		c.HTML(http.StatusOK, "timeline.tpl", gin.H{
			"title":     "My Timeline",
			"user":      user,
			"private":   true,
			"user_page": true,
			"flashes":   data,
			"messages":  GetMessages(user, page, c),
		})
	}
}