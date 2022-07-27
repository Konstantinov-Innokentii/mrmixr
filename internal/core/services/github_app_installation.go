package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type GithubAppInstallationService struct {
	githubInstallationRepo ports.GithubAppInstallationRepo
	githubInstallationAPI  ports.GithubAppInstallationAPI
}

func NewGithubInstallationService(gaiRepo ports.GithubAppInstallationRepo, gaiApi ports.GithubAppInstallationAPI) *GithubAppInstallationService {
	service := &GithubAppInstallationService{
		githubInstallationRepo: gaiRepo,
		githubInstallationAPI:  gaiApi,
	}
	return service
}

func (s *GithubAppInstallationService) GetByInstallationID(ctx context.Context, installationID int) (*domain.GithubAppInstallation, error) {
	i, err := s.githubInstallationRepo.GetByInstallationID(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return i, nil

}

func (s *GithubAppInstallationService) Store(ctx context.Context, installation *domain.GithubAppInstallation) error {
	err := s.githubInstallationRepo.InsertInstallation(ctx, installation)
	return err
}

func (s *GithubAppInstallationService) GetInstallationType(ctx context.Context, installationID int) (*domain.InstallationType, error) {
	installationType, err := s.githubInstallationAPI.GetInstallationType(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return installationType, nil
}
