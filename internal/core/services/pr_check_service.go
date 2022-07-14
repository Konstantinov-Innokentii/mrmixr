package services

import (
	"context"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/entities"
	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
)

type PrCheckService struct {
	PrCheckRepository repositories.PRCheckRepository
}

func NewPrCheckService(r repositories.PRCheckRepository) *PrCheckService {
	service := &PrCheckService{
		PrCheckRepository: r,
	}
	return service
}

func (service *PrCheckService) Store(ctx context.Context, prCheck *entities.PRCheck) (err error) {
	err = service.PrCheckRepository.Insert(ctx, prCheck)
	return err
}
