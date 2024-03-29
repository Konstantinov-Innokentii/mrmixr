package github_rest_api_v3

import (
	"context"
	entities "github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/github_api"
)

type GithubRepositoryAPI struct {
	GithubApi *github_api.GithubAPI
}

func (repoAPI GithubRepositoryAPI) List(ctx context.Context, installation *entities.GithubAppInstallation) ([]*entities.GithubRepository, error) {
	client, err := repoAPI.GithubApi.NewInstallationClient(installation.InstallationID)
	if err != nil {
		return nil, err
	}
	listRepos, _, err := client.Apps.ListRepos(ctx, nil)
	if err != nil {
		return nil, err
	}
	res := make([]*entities.GithubRepository, 0, len(listRepos.Repositories))
	for _, r := range listRepos.Repositories {
		res = append(res, &entities.GithubRepository{
			Name:                    *r.Name,
			GithubID:                *r.ID,
			GithubAppInstallationID: installation.ID,
		})
	}
	return res, nil
}
