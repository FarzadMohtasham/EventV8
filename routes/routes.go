package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(eventV8 *gin.Engine) {
	registerEventRoutes(eventV8)
}

func registerEventRoutes(eventV8 *gin.Engine) {
	eventV8.GET("/events", GetEvents)
	eventV8.GET("/events/:id", GetEvent)
	eventV8.POST("/events", CreateEvent)
	eventV8.PUT("/events/:id", UpdateEvent)
	eventV8.DELETE("/events/:id", DeleteEvent)

}
