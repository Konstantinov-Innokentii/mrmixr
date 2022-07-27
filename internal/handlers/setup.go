package handlers

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type SetupHandler struct {
	githubSetupService ports.GithubAppSetupService
}

func NewSetupHandler(githubAppSetupService ports.GithubAppSetupService) *SetupHandler {
	return &SetupHandler{
		githubAppSetupService,
	}
}

func (handler *SetupHandler) Setup(c *gin.Context) {
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

	c.Status(http.StatusCreated)
}
