package controllers

import (
	"fmt"
	"minitwit.com/devops/src/flash"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Timeline(c *gin.Context) {
	user, _ := c.Cookie("token")

	page := c.DefaultQuery("page", "0")

	if user == "" {
		c.HTML(http.StatusOK, "timeline.tpl", gin.H{
			"title": "Timeline",
			// "endpoint": "/public_timeline",
			"messages": GetMessages("", page),
		})
	} else {
		c.HTML(http.StatusOK, "timeline.tpl", gin.H{
			"title":         "Timeline",
			"user":          user,
			"user_timeline": false,
			"messages":      GetMessages("", page),
		})
	}
}

func UserTimeline(c *gin.Context) {
	user_query := c.Request.URL.Query().Get("username")
	data := make(map[string]interface{})
	data["messages"] = flash.GetFlash(c, "message")
	fmt.Println(data)
	fmt.Println(flash.GetFlash(c, "message"), flash.GetFlash(c, "error"))
	res := make(map[string]interface{})
	res["errors"] = flash.GetFlash(c, "error")
	//page := c.DefaultQuery("page", "0")

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
				"messages":      data,
			})
		} else {
			c.HTML(http.StatusOK, "timeline.tpl", gin.H{
				"title":         user_query + "'s Timeline",
				"user_timeline": true,
				"private":       true,
				"messages":      data,
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
			"messages":  data,
		})
	}
}
