package controller

import (
	"github.com/gin-gonic/gin"
	"messageBoard/dao"
	"messageBoard/service"
	"net/http"
	"strconv"
)

// Reply 被设置为全部用户可见
//（不然如果可以随便设置可见用户的话，感觉会造成逻辑上的混乱
//  比如小明回复小红，却设置为小红不可见，就很离谱；
//  或者小明发了一条评论，设为小红不可见，而小美回复了小明的这条评论，并设置回复为小红可见，也很离谱）
func ShowComments(c *gin.Context) {
	visitorComments, err := service.GetComments("")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	username, err := c.Cookie("username")
	if err != nil {
		c.JSON(http.StatusOK, visitorComments)
		return
	}
	password, err := c.Cookie("password")
	if err != nil {
		c.JSON(http.StatusOK, visitorComments)
		return
	}
	user, err := dao.GetUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	comments, err := service.GetComments(strconv.Itoa(user.Id))
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
	secretTarget := c.Query("secret_target")
	anonymous := c.Query("anonymous")
	var userId int
	if anonymous == "true" {
		userId = -1
	} else {
		userId = c.GetInt("userId")
	}
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "no value provided",
		})
		return
	}
	err := service.AddComment(value, userId, secretTarget)
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
	anonymous := c.Query("anonymous")
	var userId int
	if anonymous == "true" {
		userId = -1
	} else {
		userId = c.GetInt("userId")
	}
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
		"Info": "liked successfully",
	})
}
func DeleteComment(c *gin.Context) {
	target := c.Query("target")
	userId := c.GetInt("userId")
	if target == "" {
		c.JSON(http.StatusOK, gin.H{
			"Error": "no target provided",
		})
		return
	}
	comment, err := dao.GetCommentById(target)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
		return
	}
	user, err := dao.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err,
		})
		return
	}
	if comment.UserId != user.Id {
		c.JSON(http.StatusOK, gin.H{
			"Error": "the user does not match the comment",
		})
		return
	}
	err = service.DeleteComment(target)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Info": "deleted successfully",
	})

}
