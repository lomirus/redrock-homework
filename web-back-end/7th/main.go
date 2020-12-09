package main

import (
	"log"
	"messageBoard/controller"
)

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", controller.Root)
	router.GET("/login", controller.Login)
	router.GET("/register", controller.Register)
	router.GET("/logout", controller.Logout)

	editGroup := router.Group("edit")
	editGroup.GET("/username", controller.Verify(), controller.EditUsername)
	editGroup.GET("/password", controller.Verify(), controller.EditPassword)
	editGroup.GET("/bio", controller.Verify(), controller.EditBio)

	commentGroup := router.Group("comment")
	commentGroup.GET("/show", controller.ShowComments)
	commentGroup.GET("/add", controller.Verify(), controller.AddComment)
	commentGroup.GET("/reply", controller.Verify(), controller.ReplyComment)
	commentGroup.GET("/like", controller.Verify(), controller.LikeComment)
	commentGroup.GET("/delete", controller.Verify(), controller.DeleteComment)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
