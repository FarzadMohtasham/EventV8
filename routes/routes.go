package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginEngine *gin.Engine) {

	ginEngine.GET("/events", GetEvents)
	ginEngine.GET("/events/:id", GetEvent)
	ginEngine.POST("/events", CreateEvent)
	ginEngine.PUT("/events/:id", UpdateEvent)
}
