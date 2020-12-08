package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messageBoard/dao"
	"messageBoard/service"
	"net/http"
	"strconv"
)

func Root(c *gin.Context) {
	var user dao.User
	var err error
	user.Username, err = c.Cookie("username")
	if err != nil {
		c.String(http.StatusOK, "Please Login at '/login'")
		return
	}
	user.Password, err = c.Cookie("password")
	if err != nil {
		c.String(http.StatusOK, "Please Login at '/login'")
		return
	}
	strId, err := c.Cookie("id")
	if err != nil {
		c.String(http.StatusOK, "Please Login at '/login'")
		return
	}
	user.Id, _ = strconv.Atoi(strId)
	user.Bio, err = c.Cookie("bio")
	if err != nil {
		c.String(http.StatusOK, "Please Login at '/login'")
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s!\nYour ID is %d;\nYour Bio is %s.",
		user.Username, user.Id, user.Bio))
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Invalid Username",
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "Invalid Password",
		})
		return
	}
	user, err := service.Login(username, password)
	if err == nil {
		c.SetCookie("username", user.Username, 3600, "/", "localhost", false, true)
		c.SetCookie("password", user.Password, 3600, "/", "localhost", false, true)
		c.SetCookie("id", fmt.Sprintf("%d", user.Id), 3600, "/", "localhost", false, true)
		c.SetCookie("bio", user.Bio, 3600, "/", "localhost", false, true)
		c.String(http.StatusOK, "Login Successfully")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	}
}
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	err := service.Register(username, password)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"Info": "Registration Successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
	}
}
func Logout(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", "localhost", false, true)
	c.SetCookie("password", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"Info": "Logged out Successfully",
	})
}
