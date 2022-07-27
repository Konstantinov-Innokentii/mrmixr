package domain

type GithubRepository struct {
	ID                      int
	Name                    string
	GithubID                int64 `db:"github_id"`
	GithubAppInstallationID int   `db:"github_app_installation_id"`
}
