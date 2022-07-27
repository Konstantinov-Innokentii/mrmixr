package ports

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubRepositoryService interface {
	StoreBatch(ctx context.Context, repos []*domain.GithubRepository) error
	Fetch(ctx context.Context, installation *domain.GithubAppInstallation) error
	ListByGithubAppInstallationID(ctx context.Context, githubAppInstallationID int) ([]*domain.GithubRepository, error)
}
