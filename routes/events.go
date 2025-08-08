package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/FarzadMohtasham/EventV8/models"
	"github.com/FarzadMohtasham/EventV8/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
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

func createEvent(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Not Authorized",
		})
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Not Authorized",
		})
		return
	}

	var events models.Event
	err = ctx.ShouldBindJSON(&events)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not parse request",
		})
		return
	}

	newEvent := models.Event{
		ID:          1,
		UserID:      1,
		Name:        events.Name,
		Description: events.Description,
		Location:    events.Location,
		DateTime:    events.DateTime,
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

func getEvent(ctx *gin.Context) {
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

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not parse event id",
		})
		return
	}

	_, err = models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not handle request",
		})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not handle request",
		})
		return
	}

	updatedEvent.ID = int(eventId)
	err = updatedEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not update event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   true,
		"message": "Event updated successfully",
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not handle request",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not fetch event",
		})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not delete event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Event deleted successfully",
	})
}
