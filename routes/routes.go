package routes

import (
	"github.com/VituSuperMEg/take-tree-go/handlers"
	"github.com/gin-gonic/gin"
)

func iniatializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", handlers.ListUsers)
		v1.POST("/users", handlers.CreateUser)
	}
}
