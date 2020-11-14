package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type User struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type Dalao struct {
	Name    string
	Preview string
	QQ      int
}

var users []User
var dalaos = []Dalao{
	{
		Name:    "翔哥",
		Preview: "No Name",
		QQ:      1297880480,
	}, {
		Name:    "峰峰子",
		Preview: "sarail",
		QQ:      847117505,
	}, {
		Name:    "鑫鑫学姐",
		Preview: "小宇",
		QQ:      735268835,
	}, {
		Name:    "欣欣学姐",
		Preview: "风不经意",
		QQ:      1136843910,
	}, {
		Name:    "星星学姐",
		Preview: "浪花号",
		QQ:      3077006505,
	}}

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
	router.GET("/draw", func(c *gin.Context) {
		i := rand.Intn(5)
		c.JSON(http.StatusOK, gin.H{
			"name":    dalaos[i].Name,
			"preview": dalaos[i].Preview,
			// 下面地址中的 “&” 有时会被转义为 “\u0026”
			// 但经测试只有在 Chrome 中才会被转义，在 Postman 则显示正常
			// 所以说看来不是我的问题啦2333那就不修了（）
			"image": "https://q.qlogo.cn/headimg_dl?dst_uin=" + fmt.Sprintf("%d", dalaos[i].QQ) + "&spec=100",
		})
	})

	router.Run(":8080")
}
