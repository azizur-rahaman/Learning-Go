package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine) {
	engine.GET("/events", getEvents)
	engine.GET("/event/:id", getEventById)
	engine.POST("/events", authenticate, createEvent)
	engine.PUT("/event/:id", authenticate, updateEvent)
	engine.DELETE("/event/:id", authenticate, deleteEvent)

	engine.POST("/event/:id/register", authenticate, registerForEvent)
	engine.DELETE("/event/:id/register", authenticate, unregisterFromEvent)
	engine.GET("/event/:id/registrations", authenticate, getEventRegistrations)
	engine.GET("/event/:id/registration-status", authenticate, checkRegistrationStatus)
	engine.GET("/user/registrations", authenticate, getUserRegistrations)

	engine.POST("/signup", authLogging, inputValidation, signUp)
	engine.POST("/login", authLogging, inputValidation, loginRateLimit, login)
}
