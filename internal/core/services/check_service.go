package services

//
//import (
//	"context"
//	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/domain"
//	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/ports/repositories"
//)
//
//type CheckService struct {
//	CheckRepository repositories.CheckRepository
//}
//
//func NewCheckService(r repositories.CheckRepository) *CheckService {
//	service := &CheckService{
//		CheckRepository: r,
//	}
//	return service
//}
//
//func (service *CheckService) Store(ctx context.Context, Check *domain.Check) error {
//	err := service.CheckRepository.Insert(ctx, Check)
//	return err
//}
//
//func (service *CheckService) List(ctx context.Context) ([]*domain.Check, error) {
//	checks, err := service.CheckRepository.List(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return checks, nil
//}
//
//func (service *CheckService) RetrieveCheckByID(ctx context.Context, id int) (*domain.Check, error) {
//	check, err := service.CheckRepository.GetByID(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//	return check, nil
//}
//
//func (service *CheckService) Delete(ctx context.Context, id int) error {
//	err := service.CheckRepository.Delete(ctx, id)
//	return err
//}
