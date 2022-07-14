package entities

type GithubRepository struct {
	Id                   int
	Name                 string
	GithubId             int64 `db:"github_id"`
	GithubInstallationId int   `db:"github_installation_id"`
}
