package company

import (
	"context"
	"errors"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	companyPolicy "github.com/POMBNK/shtrafovNetTestTask/internal/domain/policy/company"
	"github.com/POMBNK/shtrafovNetTestTask/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCompany(
	ctx context.Context,
	req *protoCompanyService.GetCompanyRequest,
) (*protoCompanyService.GetCompanyResponse, error) {

	logger.L(ctx).Info("GetCompany")

	company, err := s.policy.GetCompanyByInn(ctx, req.GetInn())
	if err != nil {
		if errors.Is(err, companyPolicy.ErrInnInvalid) {
			logger.WithError(ctx, err).Error("GetCompany")
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}

	response := NewGetCompanyResponse(company)

	return response, nil
}
