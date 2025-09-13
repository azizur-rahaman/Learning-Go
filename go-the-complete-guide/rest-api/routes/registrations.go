package routes

import (
	"azizur/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// registerForEvent - Register user for an event
func registerForEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userID := ctx.GetInt64("userId")

	// Check if event exists
	_, err = models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	// Check if user is already registered
	alreadyRegistered, err := models.CheckRegistrationExists(eventID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status"})
		return
	}

	if alreadyRegistered {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Already registered for this event"})
		return
	}

	// Create registration
	registration := models.Registration{
		EventID: eventID,
		UserID:  userID,
	}

	err = registration.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully registered for event",
		"registration": gin.H{
			"eventId": eventID,
			"userId":  userID,
		},
	})
}

// unregisterFromEvent - Unregister user from an event
func unregisterFromEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userID := ctx.GetInt64("userId")

	// Check if registration exists
	registrationExists, err := models.CheckRegistrationExists(eventID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status"})
		return
	}

	if !registrationExists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not registered for this event"})
		return
	}

	// Delete registration
	registration := models.Registration{
		EventID: eventID,
		UserID:  userID,
	}

	err = registration.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}

// getEventRegistrations - Get all registrations for an event
func getEventRegistrations(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Check if event exists
	_, err = models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	registrations, err := models.GetRegistrationsByEvent(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"eventId":       eventID,
		"registrations": registrations,
		"count":         len(registrations),
	})
}

// getUserRegistrations - Get all events a user is registered for
func getUserRegistrations(ctx *gin.Context) {
	userID := ctx.GetInt64("userId")

	registrations, err := models.GetRegistrationsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user registrations"})
		return
	}

	// Get event details for each registration
	var registeredEvents []gin.H
	for _, registration := range registrations {
		event, err := models.GetEventById(registration.EventID)
		if err != nil {
			continue // Skip if event not found
		}

		registeredEvents = append(registeredEvents, gin.H{
			"registrationId": registration.ID,
			"event":          event,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"userId":           userID,
		"registeredEvents": registeredEvents,
		"count":            len(registeredEvents),
	})
}

// checkRegistrationStatus - Check if user is registered for a specific event
func checkRegistrationStatus(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userID := ctx.GetInt64("userId")

	// Check if event exists
	_, err = models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	isRegistered, err := models.CheckRegistrationExists(eventID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not check registration status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"eventId":      eventID,
		"userId":       userID,
		"isRegistered": isRegistered,
	})
}
