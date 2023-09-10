package company

import (
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/policy/company"
)

type Server struct {
	policy *company.Policy
	protoCompanyService.UnimplementedCompanyServiceServer
}

func NewServer(policy *company.Policy, srv protoCompanyService.UnimplementedCompanyServiceServer) *Server {
	return &Server{
		policy:                            policy,
		UnimplementedCompanyServiceServer: srv,
	}
}
