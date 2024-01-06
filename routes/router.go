package routes

import "github.com/gin-gonic/gin"

func InitApi() {
	r := gin.Default()
	iniatializeRoutes(r)
	r.Run()
}
