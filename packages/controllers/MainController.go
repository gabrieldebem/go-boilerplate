package controllers

import (
	"github.com/boilerplate/packages/repositories"
	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (m MainController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (m MainController) Log(c *gin.Context) {
	go repositories.LogRepository{}.Create(c.Request.UserAgent())

	c.JSON(200, gin.H{
		"message": "logged",
	})
}
