package controller

import (
	"github.com/gin-gonic/gin"
	"messageBoard/service"
	"net/http"
)

func ShowComments(c *gin.Context) {
	comments := service.ShowComments()
	c.JSON(http.StatusOK, comments)
}
func AddComment(c *gin.Context) {
	value := c.Query("value")
	userId, _ := c.Cookie("id")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Invalid Comment Value",
		})
		return
	}
	err := service.AddComment(value, userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Info": "Commented Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	}

}
func ReplyComment(c *gin.Context) {
	target := c.Query("target")
	value := c.Query("value")
	userId, _ := c.Cookie("id")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Invalid Reply Value",
		})
		return
	}
	err := service.ReplyComment(target, value, userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Info": "Replied Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	}
}
