package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type GithubAppInstallationSvc struct {
	githubInstallationRepo ports.GithubInstallationRepo
	githubInstallationAPI  ports.GithubInstallationAPI
}

func NewGithubInstallationSvc(irepo ports.GithubInstallationRepo, iapi ports.GithubInstallationAPI) *GithubAppInstallationSvc {
	svc := &GithubAppInstallationSvc{
		githubInstallationRepo: irepo,
		githubInstallationAPI:  iapi,
	}
	return svc
}

func (s *GithubAppInstallationSvc) GetByInstallationID(ctx context.Context, installationID int) (*domain.GithubAppInstallation, error) {
	i, err := s.githubInstallationRepo.GetByInstallationID(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return i, nil

}

func (s *GithubAppInstallationSvc) Store(ctx context.Context, installation *domain.GithubAppInstallation) error {
	err := s.githubInstallationRepo.InsertInstallation(ctx, installation)
	return err
}

func (s *GithubAppInstallationSvc) GetInstallationType(ctx context.Context, installationID int) (*domain.InstallationType, error) {
	installationType, err := s.githubInstallationAPI.GetInstallationType(ctx, installationID)
	if err != nil {
		return nil, err
	}
	return installationType, nil
}
