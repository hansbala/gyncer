package user

import (
	"gyncer/core"
	"gyncer/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// creates a new user
func CreateNewUserHandler(c *gin.Context) {
	var newUser database.UserCredentials
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err = database.CreateNewUserInDB(db, &newUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Failed to create new user."))
	}
}

// logs a user in and returns a new JWT token to access protected routes
func LogInUserHandler(c *gin.Context) {
	var currentUser database.UserCredentials
	if err := c.BindJSON(&currentUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "malformed input")
	}

	db, err := database.ConnectToDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to connect to database")
		return
	}

	// validates user against database
	isValidUser, err := database.IsValidUserInDB(db, &currentUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "some SQL error occured")
		return
	}
	if !isValidUser {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "you are not authorized brrrrr")
		return
	}

	// generate new JWT token
	jwtToken, err := core.GenerateJWT(currentUser.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "failed to generate JWT token")
		return
	}

	// return ok with JWT token
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
