package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FarzadMohtasham/EventV8/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	log.Println(err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not handle request",
		})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func CreateEvent(ctx *gin.Context) {
	var events models.Event
	err := ctx.ShouldBindJSON(&events)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not parse request",
		})
		return
	}

	newEvent := models.Event{
		ID:     1,
		UserID: 1,
	}

	err = newEvent.Save()

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   false,
			"message": "Could not create new event, Please try again",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "New event created successfully",
		"event":   newEvent,
	})
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Failed to process because of bad request",
		})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not process this request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  event,
	})
}
