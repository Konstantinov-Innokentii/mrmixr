package handlers

import (
	"github.com/Konstantinov-Innokentii/mrmixr/api/models"
	"github.com/Konstantinov-Innokentii/mrmixr/api/restapi/operations"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/services"
	"github.com/go-openapi/runtime/middleware"
)

type GithubRepositoriesHandler struct {
	githubRepositoriesSvc ports.GithubRepositorySvc
}

func NewGithubRepositoriesHandler(githubRepoSvc ports.GithubRepositorySvc) *GithubRepositoriesHandler {
	return &GithubRepositoriesHandler{
		githubRepositoriesSvc: githubRepoSvc,
	}
}

func (handler *GithubRepositoriesHandler) ListGithubRepositories(params operations.ListGithubRepositoriesParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()
	repos, err := handler.githubRepositoriesSvc.ListByGithubAppInstallationID(ctx, 2)
	if err != nil {
		return operations.NewListGithubRepositoriesDefault(400)
	}
	return operations.NewListGithubRepositoriesOK().WithPayload(apiGithubRepositories(repos))
}

func apiGithubRepository(in *domain.GithubRepository) *models.GithubRepository {
	return &models.GithubRepository{
		ID:   int64(in.ID),
		Name: in.Name,
	}
}

func apiGithubRepositories(in []*domain.GithubRepository) []*models.GithubRepository {
	out := make([]*models.GithubRepository, len(in))
	for i := range in {
		out[i] = apiGithubRepository(in[i])
	}
	return out
}
