package grpcServer

import v1 "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"

type Server struct {
	v1.UnimplementedCompanyServiceServer
}
