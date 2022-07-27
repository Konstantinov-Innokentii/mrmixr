package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type GithubRepositorySvc struct {
	githubRepositoryRepo ports.GithubRepositoryRepo
	githubRepositoryApi  ports.GithubRepositoryApi
}

func NewGithubRepositorySvc(grepo ports.GithubRepositoryRepo, grapi ports.GithubRepositoryApi) *GithubRepositorySvc {
	svc := &GithubRepositorySvc{
		githubRepositoryRepo: grepo,
		githubRepositoryApi:  grapi,
	}
	return svc
}

func (s *GithubRepositorySvc) StoreBatch(ctx context.Context, repos []*domain.GithubRepository) error {
	err := s.githubRepositoryRepo.InsertRepositoryBatch(ctx, repos)
	return err
}

func (s *GithubRepositorySvc) Fetch(ctx context.Context, installation *domain.GithubAppInstallation) error {
	repos, err := s.githubRepositoryApi.List(ctx, installation)
	if err != nil {
		return err
	}
	err = s.StoreBatch(ctx, repos)
	return err
}

func (s *GithubRepositorySvc) ListByGithubAppInstallationID(ctx context.Context, installationID int) ([]*domain.GithubRepository, error) {
	repos, err := s.githubRepositoryRepo.ListByGithubAppInstallationID(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return repos, nil
}
