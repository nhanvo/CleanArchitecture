package infrastructure

import (
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init()  {
	router := gin.Default()
	router.POST("/users", func(c *gin.Context) {

	})
}
