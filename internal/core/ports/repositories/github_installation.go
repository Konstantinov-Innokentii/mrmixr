package repositories

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
)

type GithubInstallationRepository interface {
	InsertInstallation(ctx context.Context, installation *entities.GithubInstallation) (err error)
	GetInstallationByInstallationId(ctx context.Context, installationID int) (*entities.GithubInstallation, error)
}

type GithubInstallationAPI interface {
	GetInstallationType(ctx context.Context, installationID int) (*entities.InstallationType, error)
}
