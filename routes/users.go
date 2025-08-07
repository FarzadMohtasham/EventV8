package routes

import (
	"fmt"
	"net/http"

	"github.com/FarzadMohtasham/EventV8/models"
	"github.com/FarzadMohtasham/EventV8/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not parse request data.",
		})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not save user.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "User created successfully",
	})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Could not parse request",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Could not authenticate user.",
		})
		return
	}

	authToken, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not authenticate user.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Login successful!",
		"token":   authToken,
	})
}
