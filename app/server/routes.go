package server

import (
	"github.com/gin-gonic/gin"
	. "github.com/susumu-iorc/ChatGPT-Personal-Assistant/app/controllers"
)

func Router() {
	r := gin.Default()
	rApi := r.Group("/api")
	rApi.GET("/hello", HelloWorldController)
	rApi.GET("/ping", PingController)
	r.Run()

}
