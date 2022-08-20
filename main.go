package main

import (
	"github.com/boilerplate/database"
	"github.com/boilerplate/packages/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	go database.Database{}.StartDb()

	r := gin.Default()

	r.GET("/ping", controllers.MainController{}.Ping)
	r.GET("/log", controllers.MainController{}.Log)

	_ = r.Run(":8080")
}
