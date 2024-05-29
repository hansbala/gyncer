package main

import (
	"github.com/hansbala/gyncer/middleware"
	"github.com/hansbala/gyncer/syncs"
	"github.com/hansbala/gyncer/user"

	"github.com/gin-gonic/gin"
)

func StartApi() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// user auth routes
	router.POST("/users", user.CreateNewUserHandler)
	router.POST("/login", user.LogInUserHandler)

	// sync routes
	router.POST("/sync", middleware.JWTTokenAuthMiddleware(), syncs.CreateSyncHandler)
	router.POST("/startsync", middleware.JWTTokenAuthMiddleware(), syncs.StartSyncsHandler)

	router.Run(":8080")
}
