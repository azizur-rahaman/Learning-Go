package routes

import (
	"azizur/rest-api/models"
	"azizur/rest-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data!", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server Error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created", "user": user})
}

func login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	tokenString, err := util.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "userId": user.ID, "token": tokenString})

}
