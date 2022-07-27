package services

import (
	"context"
	"fmt"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/services"
)

type GithubAppSetupService struct {
	installationService ports.GithubAppInstallationService
	repositoryService   ports.GithubRepositoryService
}

func NewGithubSetupService(gaiService ports.GithubAppInstallationService, grService ports.GithubRepositoryService) *GithubAppSetupService {
	service := &GithubAppSetupService{
		installationService: gaiService,
		repositoryService:   grService,
	}
	return service
}

func (s *GithubAppSetupService) Setup(ctx context.Context, installationID int) error {
	//return errors.New("some validation error")
	installationType, err := s.installationService.GetInstallationType(ctx, installationID)
	if err != nil {
		// wrap error:
		//fmw.Errof(installiion-service-error-%w)
		return fmt.Errorf("installationService.GetInstallationType: %w", err)
	}
	newInstallation, err := domain.CreateInstallation(installationID, *installationType)
	if err != nil {
		return err
	}
	err = s.installationService.Store(ctx, newInstallation)
	installation, err := s.installationService.GetByInstallationID(ctx, installationID)
	if err != nil {
		return err
	}
	err = s.repositoryService.Fetch(ctx, installation)
	go func() {
		err := s.repositoryService.Fetch(ctx, installation)
		if err != nil {
			// TODO: log smthing
		}
	}()
	return nil
}
