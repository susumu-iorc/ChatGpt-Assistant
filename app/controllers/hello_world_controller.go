package controller

import "github.com/gin-gonic/gin"

func HelloWorldController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
