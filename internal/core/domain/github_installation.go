package domain

import "time"

type InstallationType int

const (
	UserInstallation InstallationType = iota
	OrganizationInstallation
)

type GithubInstallation struct {
	ID               int
	InstallationID   int              `db:"installation_id"`
	InstalledAt      time.Time        `db:"installed_at"`
	InstallationType InstallationType `db:"installation_type"`
}

func CreateInstallation(installationID int, installationType InstallationType) (*GithubInstallation, error) {
	return &GithubInstallation{
		InstallationID:   installationID,
		InstalledAt:      time.Now().UTC(),
		InstallationType: installationType,
	}, nil
}
