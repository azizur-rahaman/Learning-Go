package routes

import (
	"azizur/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch try again"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data!", "error": err.Error()})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserId = int(userId)

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func getEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "client side Error"})
		return
	}

	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if !event.CheckOwnership(userId) {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to update this event"})
		return
	}

	var updateEvent models.Event

	err = ctx.ShouldBindJSON(&updateEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updateEvent.ID = id

	err = updateEvent.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	if !event.CheckOwnership(userId) {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Not authorized to delete this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
