package ports

import (
	"context"
)

type GithubAppSetupSvc interface {
	Setup(ctx context.Context, installationID int) error
}
