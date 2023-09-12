package company

import (
	"context"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
)

func (s *Server) GetCompany(
	ctx context.Context,
	req *protoCompanyService.GetCompanyRequest,
) (*protoCompanyService.GetCompanyResponse, error) {

	company, err := s.policy.GetCompanyByInn(ctx, req.GetInn())
	if err != nil {
		return nil, err
	}

	response := NewGetCompanyResponse(company)

	return response, nil
}
