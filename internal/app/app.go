package app

import (
	"context"
	"errors"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	grpcCompanyServer "github.com/POMBNK/shtrafovNetTestTask/internal/controller/grpc/v1/company"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/service"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/policy/company"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	grpcHostPort = ":8081"
	httpHostPort = ":8080"
)

type App struct {
	router        *httprouter.Router
	httpServer    *http.Server
	grpcServer    *grpc.Server
	CompanyServer protoCompanyService.CompanyServiceServer
}

func NewApp(ctx context.Context) App {
	router := httprouter.New()
	companyService := service.NewService()
	companyPolicy := company.NewPolicy(companyService)
	companyServer := grpcCompanyServer.NewServer(companyPolicy, protoCompanyService.UnimplementedCompanyServiceServer{})

	return App{
		router:        router,
		CompanyServer: companyServer,
	}
}

func (a *App) Start(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return a.startHTTP(ctx)
	})

	grp.Go(func() error {
		return a.startGRPC(ctx, a.CompanyServer)
	})

	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {

	listener, err := net.Listen("tcp", httpHostPort)
	if err != nil {
		log.Fatalln(err)
	}

	a.httpServer = &http.Server{
		Handler:      a.router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("server shutdown")
		default:
			log.Fatalln("failed to start server")
		}
	}

	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Fatalln("failed to shutdown server")
	}

	return err
}

func (a *App) startGRPC(ctx context.Context, server protoCompanyService.CompanyServiceServer) error {

	listener, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatalln(err)
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)

	protoCompanyService.RegisterCompanyServiceServer(a.grpcServer, server)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}
