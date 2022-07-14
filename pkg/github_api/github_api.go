package github_api

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v44/github"
	"net/http"
)

type GithubApiConfig struct {
	appId      int64
	privateKey []byte
}

func NewGithubApiConfig(appId int64, privateKey []byte) *GithubApiConfig {
	return &GithubApiConfig{
		appId:      appId,
		privateKey: privateKey,
	}
}

type GithubAPI struct {
	config    *GithubApiConfig
	transport http.RoundTripper
}

func NewGithubAPI(config *GithubApiConfig, transport http.RoundTripper) *GithubAPI {
	return &GithubAPI{
		config,
		transport,
	}
}

// NewInstallationClient create new Github API with installation token auth client for the installation.
// It is used with endpoints which requires installation token auth.
func (api *GithubAPI) NewInstallationClient(installation_id int) (*github.Client, error) {
	itr, err := ghinstallation.New(api.transport, api.config.appId, int64(installation_id), api.config.privateKey)
	if err != nil {
		return nil, err
	}
	client := github.NewClient(&http.Client{Transport: itr})

	return client, nil
}

// NewJWTClient  create new GithubAPI client  with JWT auth for the installation.
// It is used with endpoints which require JWT auth.
func (api *GithubAPI) NewJWTClient() (*github.Client, error) {
	jwtTransport, err := ghinstallation.NewAppsTransport(api.transport, api.config.appId, api.config.privateKey)
	if err != nil {
		return nil, err
	}
	client := github.NewClient(&http.Client{Transport: jwtTransport})
	return client, nil
}
