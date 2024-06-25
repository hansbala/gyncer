package youtube

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hansbala/gyncer/config"
	"github.com/hansbala/gyncer/database"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
)

const (
	// TODO: This should correspond to a unique session
	cState = "abc123"
)

var oauth2Config *oauth2.Config

func getOauth2Config() *oauth2.Config {
	if oauth2Config != nil {
		return oauth2Config
	}
	youtubeConfig := config.GetConfig().Youtube
	oauth2Config = &oauth2.Config{
		ClientID:     youtubeConfig.GyncerClientId,
		ClientSecret: youtubeConfig.GyncerClientSecret,
		RedirectURL:  youtubeConfig.GyncerRedirectUrl,
		Scopes:       []string{youtube.YoutubeScope},
	}
	return oauth2Config
}

// Generates a auth url for the user to visit.
func CreateAuthUrlHandler(c *gin.Context) {
	oauth2Config := getOauth2Config()
	authUrl := oauth2Config.AuthCodeURL(cState, oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{"auth_url": authUrl})
}

type authenticateUserRequest struct {
	UserId string `json:"user_id"`
}

// Authenticates the user and saves the credentials to the database.
func AuthenticateUserHandler(c *gin.Context) {
	// extract user id from request
	var request authenticateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "malformed request"})
		return
	}
	// get auth token from request query params
	oauth2Config := getOauth2Config()
	token, err := oauth2Config.Exchange(c.Request.Context(), c.Query("code"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "could not get token"})
		return
	}

	client := oauth2Config.Client(c.Request.Context(), token)
	service, err := youtube.New(client)
	// check if all is okay - this should be an authenticated client
	_, err = service.Channels.List([]string{"snippet"}).Do()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "could not get current user" + err.Error()})
		return
	}

	// save the credentials to database
	youtubeCredentials := database.YoutubeCredential{
		Token:  token,
		UserId: request.UserId,
	}
	db, err := database.ConnectToDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not connect to database"})
		return
	}
	if err = database.SaveYoutubeCredentials(db, youtubeCredentials); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not save credentials to database" + err.Error()})
		return
	}

	// all okay - send a 200
	c.Status(http.StatusOK)
}
