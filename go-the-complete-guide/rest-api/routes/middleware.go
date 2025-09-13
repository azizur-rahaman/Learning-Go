package routes

import (
	"azizur/rest-api/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// authenticate middleware for JWT token verification
func authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication required. Please provide a valid token.",
		})
		return
	}

	// Remove "Bearer " prefix if present
	token = strings.TrimPrefix(token, "Bearer ")

	userId, err := util.VerifyToken(token)

	if err != nil {
		var message string
		switch err.Error() {
		case "token expired":
			message = "Token has expired. Please login again."
		case "invalid token":
			message = "Invalid token. Please login again."
		default:
			message = "Authentication failed. Please provide a valid token."
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": message,
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}

// optionalAuth middleware for routes that can work with or without authentication
func optionalAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token != "" {
		token = strings.TrimPrefix(token, "Bearer ")
		userId, err := util.VerifyToken(token)
		if err == nil {
			ctx.Set("userId", userId)
			ctx.Set("authenticated", true)
		}
	}

	ctx.Next()
}
