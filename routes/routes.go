package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eventV8Engine *gin.Engine) {
	registerEventRoutes(eventV8Engine)
	registerAuthRoutes(eventV8Engine)
}

func registerEventRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.GET("/events", GetEvents)
	eventV8Engine.GET("/events/:id", GetEvent)
	eventV8Engine.POST("/events", CreateEvent)
	eventV8Engine.PUT("/events/:id", UpdateEvent)
	eventV8Engine.DELETE("/events/:id", DeleteEvent)
}

func registerAuthRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.POST("/signup", signup)
}
