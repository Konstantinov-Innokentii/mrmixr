package repositories

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
)

type CheckRepository interface {
	Insert(ctx context.Context, Check *domain.Check) (err error)
	List(ctx context.Context) ([]*domain.Check, error)
	GetByID(ctx context.Context, ID int) (*domain.Check, error)
	Delete(ctx context.Context, ID int) error
}
