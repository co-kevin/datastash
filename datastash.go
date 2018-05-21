package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hudl/fargo"
)

func main() {
	enableEurekaClient()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9999") // listen and serve on 0.0.0.0:9999
}

func enableEurekaClient() {
	c := fargo.NewConn("http://localhost:8761/eureka/v2")
	c.RegisterInstance(&fargo.Instance{
		App: "datastash",
	})
}