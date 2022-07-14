package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
)

type GithubSetupService struct {
	installationService *GithubInstallationService
	repositoryService   *GithubRepositoryService
}

func NewGithubSetupService(ii *GithubInstallationService, ri *GithubRepositoryService) *GithubSetupService {
	service := &GithubSetupService{
		installationService: ii,
		repositoryService:   ri,
	}
	return service
}

func (service *GithubSetupService) Setup(ctx context.Context, installationID int) error {
	installationType, err := service.installationService.GetInstallationType(ctx, installationID)
	if err != nil {
		return err
	}
	newInstallation, err := entities.CreateInstallation(installationID, *installationType)
	if err != nil {
		return err
	}
	err = service.installationService.Store(ctx, newInstallation)
	installation, err := service.installationService.GetByInstallationId(ctx, installationID)
	if err != nil {
		return err
	}
	err = service.repositoryService.Fetch(ctx, installation)
	go func() {
		err := service.repositoryService.Fetch(ctx, installation)
		if err != nil {
			// TODO: log smthing
		}
	}()
	return nil
}
