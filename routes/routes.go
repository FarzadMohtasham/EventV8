package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eventV8Engine *gin.Engine) {
	registerEventRoutes(eventV8Engine)
	registerAuthRoutes(eventV8Engine)
}

func registerEventRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.GET("/events", getEvents)
	eventV8Engine.GET("/events/:id", getEvent)
	eventV8Engine.POST("/events", createEvent)
	eventV8Engine.PUT("/events/:id", updateEvent)
	eventV8Engine.DELETE("/events/:id", deleteEvent)
}

func registerAuthRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.POST("/signup", signup)
	eventV8Engine.POST("/login", login)
}
