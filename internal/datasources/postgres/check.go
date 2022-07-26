package postgres

import (
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/postgres"
)

type CheckRepo struct {
	PG *postgres.Postgres
}

func NewCheckRepo(pg *postgres.Postgres) *CheckRepo {
	repo := &CheckRepo{
		PG: pg,
	}
	return repo
}

type CheckHasRepo struct {
	CheckID int `db:"check_id"`
	RepoID  int `db:"repo_id"`
}

//func (r *CheckRepo) Insert(ctx context.Context, check *domain.Check) error {
//	// TODO: execute in transaction
//	sql, args, err := r.PG.Builder.Insert("pr_check").Columns("hour", "minute", "days", "tz").Values(check.Hour, check.Minute, pq.Array(check.Days), check.Tz).Suffix("RETURNING \"id\"").ToSql()
//	if err != nil {
//		return err
//	}
//	checkID := 0
//	err = r.PG.DB.QueryRowContext(ctx, sql, args...).Scan(&checkID)
//	if err != nil {
//		return err
//	}
//
//	checkHasRepos := make([]CheckHasRepo, 0, len(*check.Repos))
//	for _, repo := range *check.Repos {
//		checkHasRepos = append(checkHasRepos, CheckHasRepo{
//			CheckID: checkID,
//			RepoID:  repo.ID,
//		})
//	}
//	// TODO: rewrite with squirell
//	_, err = r.PG.DB.NamedExec(`INSERT INTO check_has_repository (check_id, github_repository_id) VALUES (:check_id, :repo_id)`, checkHasRepos)
//	return err
//}
//
//func (r *CheckRepo) List(ctx context.Context) ([]*domain.Check, error) {
//	// TODO: execute in transaction
//	sql, args, err := r.PG.Builder.Insert("pr_check").Columns("hour", "minute", "days", "tz").Values(check.Hour, check.Minute, pq.Array(check.Days), check.Tz).Suffix("RETURNING \"id\"").ToSql()
//	if err != nil {
//		return err
//	}
//	checkID := 0
//	err = r.PG.DB.QueryRowContext(ctx, sql, args...).Scan(&checkID)
//	if err != nil {
//		return err
//	}
//
//	checkHasRepos := make([]CheckHasRepo, 0, len(*check.Repos))
//	for _, repo := range *check.Repos {
//		checkHasRepos = append(checkHasRepos, CheckHasRepo{
//			CheckID: checkID,
//			RepoID:  repo.ID,
//		})
//	}
//	// TODO: rewrite with squirell
//	_, err = r.PG.DB.NamedExec(`INSERT INTO check_has_repository (check_id, github_repository_id) VALUES (:check_id, :repo_id)`, checkHasRepos)
//	return err
//}
