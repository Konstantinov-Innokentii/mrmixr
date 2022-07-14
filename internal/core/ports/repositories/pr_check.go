package repositories

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
)

type PRCheckRepository interface {
	Insert(ctx context.Context, prCheck *entities.PRCheck) (err error)
	List(ctx context.Context) (*[]entities.PRCheck, error)
	Get(ctx context.Context) (*entities.PRCheck, error)
}
