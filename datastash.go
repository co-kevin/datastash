package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/health", health)
	r.GET("/info", info)
	r.POST("/rpc/stash", stash)
	r.Run(":" + strconv.Itoa(port)) // listen and serve on 0.0.0.0:9999
}