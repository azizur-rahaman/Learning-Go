package routes

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter stores rate limiting information for each IP
type RateLimiter struct {
	attempts map[string][]time.Time
	mutex    sync.RWMutex
}

var rateLimiter = &RateLimiter{
	attempts: make(map[string][]time.Time),
}

// loginRateLimit middleware to prevent brute force attacks on login
func loginRateLimit(ctx *gin.Context) {
	clientIP := ctx.ClientIP()

	rateLimiter.mutex.Lock()
	defer rateLimiter.mutex.Unlock()

	now := time.Now()
	windowStart := now.Add(-15 * time.Minute) // 15-minute window
	maxAttempts := 5                          // Maximum 5 attempts per 15 minutes

	// Get attempts for this IP
	attempts := rateLimiter.attempts[clientIP]

	// Remove old attempts outside the window
	var validAttempts []time.Time
	for _, attempt := range attempts {
		if attempt.After(windowStart) {
			validAttempts = append(validAttempts, attempt)
		}
	}

	// Check if limit exceeded
	if len(validAttempts) >= maxAttempts {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Too many login attempts. Please try again later.",
		})
		return
	}

	// Add current attempt
	validAttempts = append(validAttempts, now)
	rateLimiter.attempts[clientIP] = validAttempts

	ctx.Next()
}

// inputValidation middleware to validate JSON input
func inputValidation(ctx *gin.Context) {
	contentType := ctx.GetHeader("Content-Type")
	if contentType != "application/json" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Content-Type must be application/json",
		})
		return
	}

	ctx.Next()
}

// logging middleware to log authentication attempts
func authLogging(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	// Log after processing
	duration := time.Since(start)
	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	statusCode := ctx.Writer.Status()

	// You can replace this with a proper logger
	if path == "/login" || path == "/signup" {
		if statusCode >= 400 {
			// Log failed attempts
			gin.DefaultWriter.Write([]byte(
				time.Now().Format("2006/01/02 15:04:05") +
					" [AUTH-FAIL] " + clientIP + " " + method + " " + path +
					" " + string(rune(statusCode)) + " " + duration.String() + "\n"))
		} else {
			// Log successful attempts
			gin.DefaultWriter.Write([]byte(
				time.Now().Format("2006/01/02 15:04:05") +
					" [AUTH-SUCCESS] " + clientIP + " " + method + " " + path +
					" " + string(rune(statusCode)) + " " + duration.String() + "\n"))
		}
	}
}
