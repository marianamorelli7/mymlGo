package main

import (
	"github.com/gin-gonic/gin"
	controller "github.com/mercadolibre/myml/src/api/controllers/myml"
)

const (
	port= ":8080"
)

var (
	router = gin.Default()
)

func main(){
	//router.GET("/ping", ping.Ping)
	router.GET("/myml/:userID", controller.Get)
	router.Run(port)
}

