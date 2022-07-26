package github_rest_api_v3

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/github_api"
)

type GithubInstallationAPI struct {
	GithubApi *github_api.GithubAPI
}

func (repoAPI GithubInstallationAPI) GetInstallationType(ctx context.Context, installationID int) (*domain.InstallationType, error) {
	client, err := repoAPI.GithubApi.NewJWTClient()

	if err != nil {
		return nil, err
	}

	installation, _, err := client.Apps.GetInstallation(ctx, int64(installationID))

	var installationType domain.InstallationType
	switch *installation.TargetType {
	case "User":
		installationType = domain.UserInstallation
	case "Organization":
		installationType = domain.OrganizationInstallation
	}

	if err != nil {
		return nil, err
	}
	return &installationType, nil
}
