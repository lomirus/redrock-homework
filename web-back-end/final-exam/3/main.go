package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
import "card/controller"

func main() {
	var router = gin.Default()
	router.GET("/register", controller.Register())
	router.GET("/login", controller.Login())
	router.GET("/charge", controller.Charge())
	router.GET("/transfer", controller.Transfer())
	router.GET("/getLogs", controller.GetLogs())
	router.GET("/help", controller.Help())
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
