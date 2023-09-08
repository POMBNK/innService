package company_service

import (
	"context"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
)

type CompanyServer struct {
	protoCompanyService.UnimplementedCompanyServiceServer
}

func NewCompanyServer(unimplementedCompanyServiceServer protoCompanyService.UnimplementedCompanyServiceServer) *CompanyServer {
	return &CompanyServer{unimplementedCompanyServiceServer}
}

func (s *CompanyServer) GetCompany(ctx context.Context, req *protoCompanyService.GetCompanyRequest) (*protoCompanyService.GetCompanyResponse, error) {
	return nil, nil
}
