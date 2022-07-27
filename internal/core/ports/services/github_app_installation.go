package ports

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type GithubAppInstallationService interface {
	GetByInstallationID(ctx context.Context, installationID int) (*domain.GithubAppInstallation, error)
	Store(ctx context.Context, installation *domain.GithubAppInstallation) error
	GetInstallationType(ctx context.Context, id int) (*domain.InstallationType, error)
}
