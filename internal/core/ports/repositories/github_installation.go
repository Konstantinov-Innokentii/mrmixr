package ports

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubAppInstallationRepo interface {
	InsertInstallation(ctx context.Context, installation *domain.GithubAppInstallation) (err error)
	GetByInstallationID(ctx context.Context, installationID int) (*domain.GithubAppInstallation, error)
	GetInstallationByInstallationID(ctx context.Context, id int) (interface{}, interface{})
}

type GithubAppInstallationAPI interface {
	GetInstallationType(ctx context.Context, installationID int) (*domain.InstallationType, error)
}
