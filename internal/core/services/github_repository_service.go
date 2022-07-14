package services

import (
	"context"
	entities "github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type GithubRepositoryService struct {
	GithubRepositoryRepository repositories.GithubRepositoryRepository
	GithubRepositoryApi        repositories.GithubRepositoryApi
}

func NewGithubRepositoryService(grepo repositories.GithubRepositoryRepository, grapi repositories.GithubRepositoryApi) *GithubRepositoryService {
	service := &GithubRepositoryService{
		GithubRepositoryRepository: grepo,
		GithubRepositoryApi:        grapi,
	}
	return service
}

func (service *GithubRepositoryService) StoreBatch(ctx context.Context, repos *[]entities.GithubRepository) error {
	err := service.GithubRepositoryRepository.InsertRepositoryBatch(ctx, repos)
	return err
}

func (service *GithubRepositoryService) Fetch(ctx context.Context, installation *entities.GithubInstallation) error {
	repos, err := service.GithubRepositoryApi.List(ctx, installation)
	if err != nil {
		return err
	}
	err = service.StoreBatch(ctx, repos)
	return err
}
