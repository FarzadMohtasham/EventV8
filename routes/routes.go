package routes

import (
	"github.com/FarzadMohtasham/EventV8/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eventV8Engine *gin.Engine) {
	registerEventRoutes(eventV8Engine)
	registerAuthRoutes(eventV8Engine)
}

func registerEventRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.GET("/events", getEvents)
	eventV8Engine.GET("/events/:id", getEvent)

	authenticated := eventV8Engine.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}

func registerAuthRoutes(eventV8Engine *gin.Engine) {
	eventV8Engine.POST("/signup", signup)
	eventV8Engine.POST("/login", login)
}
