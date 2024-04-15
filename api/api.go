package main

import (
	"gyncer/middleware"
	"gyncer/syncs"
	"gyncer/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// user auth routes
	router.POST("/users", user.CreateNewUserHandler)
	router.POST("/login", user.LogInUserHandler)

	// sync routes
	router.POST("/sync", middleware.JWTTokenAuthMiddleware(), syncs.CreateSyncHandler)

	router.Run(":8080")
}
