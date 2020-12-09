package controller

import (
	"github.com/gin-gonic/gin"
	"messageBoard/service"
	"net/http"
)

func ShowComments(c *gin.Context) {
	comments, err := service.GetComments()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, comments)
}
func AddComment(c *gin.Context) {
	value := c.Query("value")
	userId := c.GetInt("userId")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "no value provided",
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
			"Error": err.Error(),
		})
	}
}
func ReplyComment(c *gin.Context) {
	target := c.Query("target")
	value := c.Query("value")
	userId := c.GetInt("userId")
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "no value provided",
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
			"Error": err.Error(),
		})
	}
}
func LikeComment(c *gin.Context) {
	target := c.Query("target")
	if target == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "no target provided",
		})
		return
	}
	err := service.LikeComment(target)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Info": "to like successfully",
	})
}
