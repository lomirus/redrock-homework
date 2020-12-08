package controller

import (
	"github.com/gin-gonic/gin"
	"messageBoard/service"
	"net/http"
)

func EditUsername(c *gin.Context) {
	username, _ := c.Cookie("username")
	password, _ := c.Cookie("password")
	newUsername := c.Query("new")
	err := service.EditUsername(username, password, newUsername)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	} else {
		c.SetCookie("username", newUsername, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Username Successfully",
		})
	}
}
func EditPassword(c *gin.Context) {
	username, _ := c.Cookie("username")
	password, _ := c.Cookie("password")
	newPassword := c.Query("new")
	err := service.EditPassword(username, password, newPassword)
	if err == nil {
		c.SetCookie("password", newPassword, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Password Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Warning": err,
		})
	}

}
func EditBio(c *gin.Context) {
	username, _ := c.Cookie("username")
	password, _ := c.Cookie("password")
	newBio := c.Query("new")
	err := service.EditBio(username, password, newBio)
	if err == nil {
		c.SetCookie("bio", newBio, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"Info": "Updated Bio Successfully",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	}
}
