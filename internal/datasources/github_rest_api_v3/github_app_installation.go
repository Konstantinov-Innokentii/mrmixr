package github_rest_api_v3

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/github_api"
)

type GithubInstallationAPI struct {
	GithubApi *github_api.GithubAPI
}

func (repoAPI GithubInstallationAPI) GetInstallationType(ctx context.Context, installationId int) (*entities.InstallationType, error) {
	client, err := repoAPI.GithubApi.NewJWTClient()

	if err != nil {
		return nil, err
	}

	installation, _, err := client.Apps.GetInstallation(ctx, int64(installationId))

	var installationType entities.InstallationType
	switch *installation.TargetType {
	case "User":
		installationType = entities.UserInstallation
	case "Organization":
		installationType = entities.OrganizationInstallation
	}

	if err != nil {
		return nil, err
	}
	return &installationType, nil
}
