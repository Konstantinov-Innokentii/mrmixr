package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type GithubInstallationService struct {
	GithubInstallationRepository repositories.GithubInstallationRepository
	GithubInstallationAPI        repositories.GithubInstallationAPI
}

func NewGithubInstallationService(irepo repositories.GithubInstallationRepository, iapi repositories.GithubInstallationAPI) *GithubInstallationService {
	service := &GithubInstallationService{
		GithubInstallationRepository: irepo,
		GithubInstallationAPI:        iapi,
	}
	return service
}

func (service *GithubInstallationService) GetByInstallationId(ctx context.Context, installationID int) (*entities.GithubInstallation, error) {
	i, err := service.GithubInstallationRepository.GetInstallationByInstallationId(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return i, nil

}

func (service *GithubInstallationService) Store(ctx context.Context, installation *entities.GithubInstallation) error {
	err := service.GithubInstallationRepository.InsertInstallation(ctx, installation)
	return err
}

func (service *GithubInstallationService) GetInstallationType(ctx context.Context, installationID int) (*entities.InstallationType, error) {
	installationType, err := service.GithubInstallationAPI.GetInstallationType(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return installationType, nil
}
