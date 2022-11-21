package main

import (
	"ginstudy/app/content"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func init() {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/content/:id", content.Get)
	router.DELETE("/content/:id", content.DELETE)
	router.PUT("/content/:id", content.PUT)
	router.POST("/content", content.POST)
}
