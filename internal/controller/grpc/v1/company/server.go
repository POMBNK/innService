package company

import protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"

type Server struct {
	protoCompanyService.UnimplementedCompanyServiceServer
}

func NewServer(srv protoCompanyService.UnimplementedCompanyServiceServer) *Server {
	return &Server{
		UnimplementedCompanyServiceServer: srv,
	}
}
