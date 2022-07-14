package transport

import (
	"context"
	"fmt"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/services"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
	"net/http"
	"strconv"
)

type GithubOauthHandler struct {
	githubSetupService *services.GithubSetupService
}

func NewGithubOauthHandler(githubSetupService *services.GithubSetupService) *GithubOauthHandler {
	return &GithubOauthHandler{
		githubSetupService: githubSetupService,
	}
}

var (
	oauthConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		// select level of access you want https://developer.github.com/v3/oauth/#scopes
		//Scopes:   []string{"user:email", "repo"},
		Endpoint: githuboauth.Endpoint,
	}
)

func (handler *GithubOauthHandler) Auth(c *gin.Context) {
	code := c.Query("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user.Login)

	installationID, err := strconv.Atoi(c.Query("installation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid installation_id",
		})
	}
	err = handler.githubSetupService.Setup(context.Background(), installationID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}
