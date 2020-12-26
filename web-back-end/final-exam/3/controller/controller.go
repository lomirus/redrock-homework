package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)
import "card/model"

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "username cannot be nil",
			})
			return
		}
		if len(username) > 32 {
			c.JSON(http.StatusOK, gin.H{
				"error": "username is too long",
			})
			return
		}
		if password == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "password cannot be nil",
			})
			return
		}
		if len(password) > 32 {
			c.JSON(http.StatusOK, gin.H{
				"error": "password is too long",
			})
			return
		}
		err := model.Register(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"info": "registered successfully",
			})
		}
	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		if username == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "username cannot be nil",
			})
			return
		}
		if password == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "password cannot be nil",
			})
			return
		}
		err := model.Login(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.SetCookie("username", username, 365*86400, "/", "localhost", false, true)
			c.SetCookie("password", password, 365*86400, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"info": "logged in successfully",
			})
		}
	}
}
func Charge() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		password, err := c.Cookie("password")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		err = model.Login(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		money, err := strconv.Atoi(c.Query("money"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid money",
			})
			return
		}
		if money <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid money",
			})
			return
		}
		self, err := model.GetUser(username)
		err = model.Charge(self, money)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"info": "charged successfully",
			})
		}
	}
}
func Transfer() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		password, err := c.Cookie("password")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		err = model.Login(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		targetUsername := c.Query("target")
		if targetUsername == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid target",
			})
			return
		}
		money, err := strconv.Atoi(c.Query("money"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid money",
			})
			return
		}
		if money <= 0 {
			c.JSON(http.StatusOK, gin.H{
				"error": "invalid money",
			})
			return
		}
		remark := c.Query("remark")
		self, err := model.GetUser(username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "cannot find your username",
			})
			return
		}
		target, err := model.GetUser(targetUsername)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "cannot find the target",
			})
			return
		}
		if self.Money < money {
			c.JSON(http.StatusOK, gin.H{
				"error": "your money is not enough",
			})
			return
		}
		err = model.Transfer(self, target, money, remark)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"info": "transferred successfully",
			})
		}
	}
}
func Logs() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		password, err := c.Cookie("password")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "please login first",
			})
			return
		}
		err = model.Login(username, password)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		search := c.Query("search")
		logs, err := model.GetLogs(username, search)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, logs)
		}
	}
}
func Help() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"注册":   "/register",
			"登录":   "/login",
			"充值":   "/charge",
			"转账":   "/transfer",
			"转账备注": "/transfer?remark",
			"查询记录": "/logs",
			"模糊查询": "/logs?search",
			"帮助":   "/help",
		})
	}
}
