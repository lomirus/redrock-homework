package controller

import (
	"github.com/gin-gonic/gin"
	"messageBoard/service"
	"net/http"
)

func EditUsername(c *gin.Context) {
	userId := c.GetInt("userId")
	newUsername := c.Query("new")
	err := service.EditUsername(userId, newUsername)
	if err == nil {
		c.SetCookie("username", newUsername, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Username Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
	}
}
func EditPassword(c *gin.Context) {
	userId := c.GetInt("userId")
	newPassword := c.Query("new")
	err := service.EditPassword(userId, newPassword)
	if err == nil {
		c.SetCookie("password", newPassword, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Password Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Warning": err.Error(),
		})
	}
}
func EditBio(c *gin.Context) {
	userId := c.GetInt("userId")
	newBio := c.Query("new")
	err := service.EditBio(userId, newBio)
	if err == nil {
		c.SetCookie("bio", newBio, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Bio Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
	}
}
