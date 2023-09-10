package company

import (
	"context"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
)

func (s *Server) GetCompany(
	ctx context.Context,
	req *protoCompanyService.GetCompanyRequest,
) (*protoCompanyService.GetCompanyResponse, error) {

	return nil, nil
}
