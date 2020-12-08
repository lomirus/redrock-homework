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
	editGroup.GET("/username", controller.EditUsername)
	editGroup.GET("/password", controller.EditPassword)
	editGroup.GET("/bio", controller.EditBio)

	commentGroup := router.Group("comment")
	commentGroup.GET("/show", controller.ShowComments)
	commentGroup.GET("/add", controller.AddComment)
	commentGroup.GET("/reply", controller.ReplyComment)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
