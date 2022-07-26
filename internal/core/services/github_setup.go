package services

import (
	"context"
	"fmt"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
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
	//return errors.New("some validation error")
	installationType, err := service.installationService.GetInstallationType(ctx, installationID)
	if err != nil {
		// wrap error:
		//fmw.Errof(installiion-service-error-%w)
		return fmt.Errorf("installationService.GetInstallationType: %w", err)
	}
	newInstallation, err := domain.CreateInstallation(installationID, *installationType)
	if err != nil {
		return err
	}
	err = service.installationService.Store(ctx, newInstallation)
	installation, err := service.installationService.GetByInstallationID(ctx, installationID)
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
