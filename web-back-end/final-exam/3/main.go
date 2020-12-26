package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
import "card/controller"

func main() {
	var router = gin.Default()
	router.Static("/statics", "view/statics/")
	router.LoadHTMLGlob("view/template/*")
	router.GET("/", controller.Help())
	router.GET("/register", controller.Register())
	router.GET("/login", controller.Login())
	router.GET("/charge", controller.Charge())
	router.GET("/transfer", controller.Transfer())
	router.GET("/logs", controller.Logs())
	router.GET("/help", controller.Help())
	router.GET("/qrcode/transfer", controller.QrCodeTransfer())
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
