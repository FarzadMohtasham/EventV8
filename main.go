package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	serverEngine := gin.Default()

	serverEngine.GET("/events", getEvents)

	serverEngine.Run(":8880")
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"id": 12,
	})
}
