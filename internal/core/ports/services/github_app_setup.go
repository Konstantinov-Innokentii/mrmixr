package ports

import (
	"context"
)

type GithubAppSetupService interface {
	Setup(ctx context.Context, installationID int) error
}
