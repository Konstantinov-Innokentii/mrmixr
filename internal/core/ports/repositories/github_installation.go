package repositories

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubInstallationRepository interface {
	InsertInstallation(ctx context.Context, installation *domain.GithubInstallation) (err error)
	GetInstallationByInstallationID(ctx context.Context, installationID int) (*domain.GithubInstallation, error)
}

type GithubInstallationAPI interface {
	GetInstallationType(ctx context.Context, installationID int) (*domain.InstallationType, error)
}
