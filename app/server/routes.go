package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/controllers"
)

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		}
	}
}

func Router() {
	r := gin.Default()
	r.Use(errorMiddleware())
	rApi := r.Group("/api")
	rApi.GET("/hello", HelloWorldController)
	rApi.GET("/ping", PingController)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
