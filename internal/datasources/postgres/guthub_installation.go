package postgres

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
	"github.com/Konstantinov-Innokentii/mrmixr/pkg/postgres"
)

type GithubInstallationRepo struct {
	PG *postgres.Postgres
}

func NewGithubInstallationRepo(pg *postgres.Postgres) *GithubInstallationRepo {
	repo := &GithubInstallationRepo{
		PG: pg,
	}
	return repo
}

func (r *GithubInstallationRepo) GetByInstallationID(ctx context.Context, installationID int) (i *domain.GithubAppInstallation, err error) {
	sql, args, err := r.PG.Builder.Select("*").From("github_installation").Where("installation_id IN (?)", installationID).ToSql()
	if err != nil {
		return nil, err
	}
	installation := &domain.GithubAppInstallation{}

	err = r.PG.DB.QueryRowxContext(ctx, sql, args...).StructScan(installation)
	if err != nil {
		return nil, err
	}
	return installation, nil
}

func (r *GithubInstallationRepo) InsertInstallation(ctx context.Context, installation *domain.GithubAppInstallation) (err error) {
	sql, args, err := r.PG.Builder.Insert("github_installation").Columns("installation_id", "installed_at", "installation_type").Values(installation.InstallationID, installation.InstalledAt, installation.InstallationType).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		return err
	}
	id := 0
	err = r.PG.DB.QueryRowContext(ctx, sql, args...).Scan(&id)
	return err
}
