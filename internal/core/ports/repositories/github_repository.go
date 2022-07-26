package repositories

import (
	"context"
	entities "github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubRepositoryRepository interface {
	InsertRepositoryBatch(ctx context.Context, repos []*entities.GithubRepository) error
	ListByInstallationID(ctx context.Context, installationID int) ([]*entities.GithubRepository, error)
}

type GithubRepositoryApi interface {
	List(ctx context.Context, installation *entities.GithubInstallation) ([]*entities.GithubRepository, error)
}
