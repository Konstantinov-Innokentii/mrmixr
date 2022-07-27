package services

import (
	"context"
	"fmt"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/services"
)

type GithubAppSetupSvc struct {
	installationSvc ports.GithubAppInstallationSvc
	repositorySvc   ports.GithubRepositorySvc
}

func NewGithubSetupSvc(gaiSvc ports.GithubAppInstallationSvc, grSvc ports.GithubRepositorySvc) *GithubAppSetupSvc {
	svc := &GithubAppSetupSvc{
		installationSvc: gaiSvc,
		repositorySvc:   grSvc,
	}
	return svc
}

func (s *GithubAppSetupSvc) Setup(ctx context.Context, installationID int) error {
	//return errors.New("some validation error")
	installationType, err := s.installationSvc.GetInstallationType(ctx, installationID)
	if err != nil {
		// wrap error:
		//fmw.Errof(installiion-service-error-%w)
		return fmt.Errorf("installationService.GetInstallationType: %w", err)
	}
	newInstallation, err := domain.CreateInstallation(installationID, *installationType)
	if err != nil {
		return err
	}
	err = s.installationSvc.Store(ctx, newInstallation)
	installation, err := s.installationSvc.GetByInstallationID(ctx, installationID)
	if err != nil {
		return err
	}
	err = s.repositorySvc.Fetch(ctx, installation)
	go func() {
		err := s.repositorySvc.Fetch(ctx, installation)
		if err != nil {
			// TODO: log smthing
		}
	}()
	return nil
}
