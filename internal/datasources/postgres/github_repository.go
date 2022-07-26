package postgres

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
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

func (r *GithubRepositoryRepo) InsertRepositoryBatch(ctx context.Context, repos []*domain.GithubRepository) error {
	// TODO: Rewrite with squirrel
	_, err := r.PG.DB.NamedExec(`INSERT INTO github_repository (name, github_id, github_installation_id) VALUES (:name,:github_id, :github_installation_id)`, repos)
	return err
}

func (r *GithubRepositoryRepo) ListByInstallationID(ctx context.Context, installationID int) ([]*domain.GithubRepository, error) {
	// TODO: Rewrite with squirrel
	//sql, args, err := r.PG.Builder.Select("*").From("github_repositories").Where("installation_id IN (?)", installationID).ToSql()
	//if err != nil {
	//	return nil, err
	//}
	var repos []*domain.GithubRepository
	err := r.PG.DB.SelectContext(ctx, &repos, "SELECT * FROM github_repository")
	if err != nil {
		return nil, err
	}

	return repos, nil
}
