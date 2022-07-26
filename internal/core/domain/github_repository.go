package domain

type GithubRepository struct {
	ID                   int
	Name                 string
	GithubID             int64 `db:"github_id"`
	GithubInstallationID int   `db:"github_installation_id"`
}
