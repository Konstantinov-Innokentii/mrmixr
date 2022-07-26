package main

import (
	"github.com/Konstantinov-Innokentii/mrmixr/api/restapi"
	"github.com/Konstantinov-Innokentii/mrmixr/api/restapi/operations"
	config "github.com/Konstantinov-Innokentii/mrmixr/internal/config"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/services"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/datasources/github_rest_api_v3"
	postgresRepo "github.com/Konstantinov-Innokentii/mrmixr/internal/datasources/postgres"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/transport"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/github_api"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/postgres"
	"github.com/go-openapi/loads"
	"log"
	"net/http"
)

func main() {
	// load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	// Configure Postgres
	pg, err := postgres.New(cfg.PGDBName, cfg.PGUsername, cfg.PGPassword)
	if err != nil {
		log.Fatalf("Postgres is not available")
	}

	ghConfig := github_api.NewGithubApiConfig(
		cfg.GHAppID,
		cfg.GHPrivateKey,
	)

	ghApi := github_api.NewGithubAPI(
		ghConfig,
		http.DefaultTransport,
	)
	githubReposotoryApi := github_rest_api_v3.GithubRepositoryAPI{
		ghApi,
	}
	//githubInstallationApi := github_rest_api_v3.GithubInstallationAPI{
	//	ghApi,
	//}

	//githubInstallationRepo := postgresRepo.NewGithubInstallationRepo(pg)
	githubRepositoryRepo := postgresRepo.NewGithubRepositoryRepo(pg)

	githubRepositoryService := services.NewGithubRepositoryService(githubRepositoryRepo, githubReposotoryApi)
	//githubInstallationService := services.NewGithubInstallationService(githubInstallationRepo, githubInstallationApi)
	//githubSetupService := services.NewGithubSetupService(
	//	githubInstallationService,
	//	githubRepositoryService,
	//)

	//ghOauthHttpHandler := transport.NewGithubOauthHandler(githubSetupService)
	//checksHandler := transport.NewChecksHandler()
	githubRepositoryHandler := transport.NewGithubRepositoriesHandler(
		githubRepositoryService,
	)

	// TODO: move server initialization in separate package
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMrmixrAPI(swaggerSpec)

	api.ListGithubRepositoriesHandler = operations.ListGithubRepositoriesHandlerFunc(githubRepositoryHandler.ListGithubRepositories)

	//api.GetPrChecksHandler = op.GetPrChecksHandlerFunc()
	server := restapi.NewServer(api)
	server.Port = 8000
	defer server.Shutdown()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
