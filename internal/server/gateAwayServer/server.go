package gateAwayServer

import (
	"context"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	"github.com/POMBNK/shtrafovNetTestTask/internal/controller/grpc/v1/company_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const (
	grpcHostPort = ":8081"
	httpHostPort = ":8080"
)

type Server struct {
	protoCompanyService.UnimplementedCompanyServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	protoCompanyService.RegisterCompanyServiceServer(
		grpcServer,
		company_service.NewCompanyServer(protoCompanyService.UnimplementedCompanyServiceServer{}),
	)

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = protoCompanyService.RegisterCompanyServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		log.Fatalf("failed to register endpoint: %v", err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(":8081", mux)
	})

	err = g.Wait()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
