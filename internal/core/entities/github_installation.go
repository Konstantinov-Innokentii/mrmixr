package entities

import "time"

type InstallationType int

const (
	UserInstallation InstallationType = iota
	OrganizationInstallation
)

type GithubInstallation struct {
	Id               int
	InstallationId   int              `db:"installation_id"`
	InstalledAt      time.Time        `db:"installed_at"`
	InstallationType InstallationType `db:"installation_type"`
}

func CreateInstallation(installationId int, installationType InstallationType) (*GithubInstallation, error) {
	return &GithubInstallation{
		InstallationId:   installationId,
		InstalledAt:      time.Now().UTC(),
		InstallationType: installationType,
	}, nil
}
