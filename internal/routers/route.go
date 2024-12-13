package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", InitServer)
		v1.GET("/ping/:name", InitServer)
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/ping", InitServer)
	}
	return r
}
func InitServer(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}
