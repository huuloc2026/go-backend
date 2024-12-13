package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huuloc2026/go-backend.git/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/1", controllers.NewUserController().GetUserbyID)
	}
	// v2 := r.Group("/v2")
	// {
	// 	v2.GET("/ping", InitServer)
	// }
	return r
}
func InitServer(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}
