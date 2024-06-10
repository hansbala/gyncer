package spotify

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hansbala/gyncer/config"
	"github.com/hansbala/gyncer/core"
	"github.com/hansbala/gyncer/database"
	"github.com/hansbala/gyncer/middleware"
	"github.com/zmb3/spotify/v2"
	spotify_auth "github.com/zmb3/spotify/v2/auth"
)

const (
	// TODO: This should correspond to a unique session
	cState = "abc123"
)

var cAuthScopes = []string{spotify_auth.ScopePlaylistModifyPrivate, spotify_auth.ScopeUserLibraryModify}

var spotifyAuthenticator *spotify_auth.Authenticator

func getSpotifyAuthenticator() spotify_auth.Authenticator {
	if spotifyAuthenticator != nil {
		return *spotifyAuthenticator
	}
	spotifyConfig := config.GetConfig().Spotify
	spotifyAuthenticator = spotify_auth.New(
		spotify_auth.WithRedirectURL(spotifyConfig.GyncerRedirectUrl),
		spotify_auth.WithScopes(cAuthScopes...),
		spotify_auth.WithClientID(spotifyConfig.GyncerClientId),
		spotify_auth.WithClientSecret(spotifyConfig.GyncerClientSecret),
	)
	return *spotifyAuthenticator
}

// generates a auth url for the user to visit
func CreateAuthUrlHandler(c *gin.Context) {
	authUrl := getSpotifyAuthenticator().AuthURL(cState)
	c.JSON(http.StatusOK, gin.H{"auth_url": authUrl})
}

func AuthenticateUserHandler(c *gin.Context) {
	// extract user id from request
	jwtToken, err := middleware.GetTokenFromRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// get user email
	userEmail, err := core.GetEmailFromJWT(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	// get user id by hashing the email
	userId, err := core.HashString(userEmail)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "failed to hash user email"})
		return
	}

	// get auth token from request query params
	spotifyAuthenticator := getSpotifyAuthenticator()
	token, err := spotifyAuthenticator.Token(c, cState, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "could not get token"})
		return
	}

	// check if all is okay - this should be an authenticated client
	client := spotify.New(spotifyAuthenticator.Client(c, token))
	_, err = client.CurrentUser(context.TODO())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "could not get current user" + err.Error()})
		return
	}

	// save the credentials to database
	spotifyCredentials := database.SpotifyCredential{
		Token:  token,
		UserId: userId,
	}
	db, err := database.ConnectToDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not connect to database"})
		return
	}
	if err = database.SaveSpotifyCredentials(db, spotifyCredentials); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not save credentials to database" + err.Error()})
		return
	}

	// all okay - send a 200
	c.Status(http.StatusOK)
}
