package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var users []User

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		cookie, err := c.Request.Cookie("username")
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "Hello, " + cookie.Value + "!",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    200,
				"message": "Hello, visitor!",
			})
		}
	})
	router.GET("/register", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBind(&newUser); err == nil {
			var hasRegistered bool
			for i := range users {
				if users[i].Username == newUser.Username {
					hasRegistered = true
					break
				}
			}
			if hasRegistered {
				c.String(http.StatusBadRequest, "The username has been registered.")
			} else {
				users = append(users, newUser)
				c.String(http.StatusOK, "Registered successfully.")
			}
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
	})
	router.GET("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err == nil {
			var status int = 0
			//0 no user
			//1 wrong password
			//2 success
			for i := range users {
				if users[i].Username == user.Username {
					if users[i].Password == user.Password {
						status = 2
						break
					} else {
						status = 1
						break
					}
				}
			}
			switch status {
			case 0:
				c.String(http.StatusBadRequest, "The username does not exist.")
			case 1:
				c.String(http.StatusBadRequest, "Wrong password.")
			case 2:
				cookie := &http.Cookie{
					Name:     "username",
					Value:    user.Username,
					Path:     "/",
					HttpOnly: true,
				}
				http.SetCookie(c.Writer, cookie)
				c.String(http.StatusOK, "Login successful.")
			}
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
	})

	router.Run(":8080")
}
