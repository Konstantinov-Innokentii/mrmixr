package ports

import (
	"context"
	entities "github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubRepositoryRepo interface {
	InsertRepositoryBatch(ctx context.Context, repos []*entities.GithubRepository) error
	ListByGithubAppInstallationID(ctx context.Context, installationID int) ([]*entities.GithubRepository, error)
}

type GithubRepositoryApi interface {
	List(ctx context.Context, installation *entities.GithubAppInstallation) ([]*entities.GithubRepository, error)
}
