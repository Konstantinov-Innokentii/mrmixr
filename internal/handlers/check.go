package handlers

//import (
//	"github.com/Konstantinov-Innokentii/mrmixr/internal/core/services"
//)
//
//type ChecksHandler struct {
//	checkService *prt.CheckService
//}

//func apiCheck(in *domain.Check) *models.Check {
//	return &models.Check{
//		Days:               in.Days,
//		GithubRepositories: nil,
//		Hour:               int64(in.Hour),
//		ID:                 int64(in.ID),
//		Minute:             int64(in.Minute),
//		Name:               in.Name,
//		Tz:                 in.Tz,
//	}
//}
//
//func apiChecks(in []*domain.Check) []*models.Check {
//	out := make([]*models.Check, len(in))
//	for i := range in {
//		out[i] = apiCheck(in[i])
//	}
//	return out
//}
//
//func NewChecksHandler(checkSvc *services.CheckService) *ChecksHandler {
//	return &ChecksHandler{
//		checkService: checkSvc,
//	}
//}
//
//func (handler *ChecksHandler) ListChecks(params operations.ListChecksParams) operations.ListChecksResponder {
//	ctx := params.HTTPRequest.Context()
//	checks, err := handler.checkService.List(ctx)
//	if err != nil {
//		return
//	}
//	return operations.NewListChecksOK().WithPayload(apiChecks(checks))
//}
