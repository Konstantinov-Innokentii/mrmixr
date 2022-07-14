package postgres

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/postgres"
	"github.com/lib/pq"
)

type PRCheckRepo struct {
	PG *postgres.Postgres
}

func NewPRCheckRepo(pg *postgres.Postgres) *PRCheckRepo {
	repo := &PRCheckRepo{
		PG: pg,
	}
	return repo
}

type CheckHasRepo struct {
	CheckId int `db:"check_id"`
	RepoId  int `db:"repo_id"`
}

func (r *PRCheckRepo) InsertCheck(ctx context.Context, check *entities.PRCheck) error {
	// TODO: execute in transaction
	sql, args, err := r.PG.Builder.Insert("pr_check").Columns("hour", "minute", "days", "tz").Values(check.Hour, check.Minute, pq.Array(check.Days), check.Tz).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		return err
	}
	id := 0
	err = r.PG.DB.QueryRowContext(ctx, sql, args...).Scan(&id)
	if err != nil {
		return err
	}

	checkHasRepos := make([]CheckHasRepo, 0, len(*check.Repos))
	for _, repo := range *check.Repos {
		checkHasRepos = append(checkHasRepos, CheckHasRepo{
			CheckId: id,
			RepoId:  repo.Id,
		})
	}
	// TODO: rewrite with squirell
	_, err = r.PG.DB.NamedExec(`INSERT INTO check_has_repository (check_id, github_repository_id) VALUES (:check_id, :repo_id)`, checkHasRepos)
	return err
}
