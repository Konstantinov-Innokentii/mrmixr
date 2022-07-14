package postgres

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/postgres"
)

type GithubRepositoryRepo struct {
	PG *postgres.Postgres
}

func NewGithubRepositoryRepo(pg *postgres.Postgres) *GithubRepositoryRepo {
	repo := &GithubRepositoryRepo{
		PG: pg,
	}
	return repo
}

func (rr *GithubRepositoryRepo) InsertRepositoryBatch(ctx context.Context, repos *[]entities.GithubRepository) error {
	// TODO: Rewrite with squirrel?
	_, err := rr.PG.DB.NamedExec(`INSERT INTO github_repository (name, github_id, github_installation_id) VALUES (:name,:github_id, :github_installation_id)`, *repos)
	return err
}
