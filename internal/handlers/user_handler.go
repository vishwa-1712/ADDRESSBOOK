package handlers

import (
	"addressbook/internal/models"
	"addressbook/internal/services"
	"addressbook/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	//Get email/password off req body
	//Hash password
	//Create User
	// Respond

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// if !services.AuthenticateUser(input.Email, input.Password) {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }
	// Here we store the error returned from AuthenticateUser
	if err := services.AuthenticateUser(input.Email, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token gerneration failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
