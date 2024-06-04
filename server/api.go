package main

import (
	"github.com/hansbala/gyncer/middleware"
	"github.com/hansbala/gyncer/spotify"
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

	// spotify routes
	router.POST("/spotify/authurl", middleware.JWTTokenAuthMiddleware(), spotify.CreateAuthUrlHandler)
	router.POST("/spotify/callback", middleware.JWTTokenAuthMiddleware(), spotify.AuthenticateUserHandler)

	router.Run(":8080")
}
