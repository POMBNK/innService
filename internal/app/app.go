package app

import (
	"context"
	"errors"
	"fmt"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	grpcCompanyServer "github.com/POMBNK/shtrafovNetTestTask/internal/controller/grpc/v1/company"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/service"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/policy/company"
	"github.com/POMBNK/shtrafovNetTestTask/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	loopback         = "127.0.0.1"
	grpcHostPort     = "8082"
	httpGateAwayPort = "8081"
	httpHostPort     = "8080"
)

type App struct {
	router        *httprouter.Router
	httpServer    *http.Server
	grpcServer    *grpc.Server
	companyServer protoCompanyService.CompanyServiceServer
}

func NewApp(ctx context.Context) App {
	logger.L(ctx).Info("Router initialization...")
	router := httprouter.New()

	companyService := service.NewService()
	companyPolicy := company.NewPolicy(companyService)
	companyServer := grpcCompanyServer.NewServer(companyPolicy, protoCompanyService.UnimplementedCompanyServiceServer{})

	return App{
		router:        router,
		companyServer: companyServer,
	}
}

func (a *App) Start(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)

	logger.L(ctx).Info("Starting http server...")
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})

	logger.L(ctx).Info("Starting grpc server...")
	grp.Go(func() error {
		return a.startGRPC(ctx, a.companyServer)
	})
	grp.Go(func() error {
		return a.startHTTPGateAway(ctx)
	})

	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", loopback, httpHostPort))
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
			logger.L(ctx).Warn("http server closed")
		default:
			logger.WithError(ctx, err).Fatal("failed to start server")
		}
	}

	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.WithError(ctx, err).Fatal("failed to shutdown server")
	}

	return err
}
func (a *App) startHTTPGateAway(ctx context.Context) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", loopback, httpGateAwayPort))
	if err != nil {
		logger.WithError(ctx, err).Fatal("failed to create listener")
	}

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = protoCompanyService.RegisterCompanyServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", loopback, grpcHostPort), opts)
	if err != nil {
		logger.WithError(ctx, err).Fatal("failed to register handler from endpoint grpc gateway")
	}

	a.httpServer = &http.Server{
		Handler:      mux,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.L(ctx).Warn("server shutdown")
		default:
			logger.WithError(ctx, err).Fatal("failed to start server")
		}
	}

	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.WithError(ctx, err).Fatal("failed to shutdown server")
	}

	return err
}

func (a *App) startGRPC(ctx context.Context, server protoCompanyService.CompanyServiceServer) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", loopback, grpcHostPort))
	if err != nil {
		logger.WithError(ctx, err).Fatal("failed to create listener")
	}

	var serverOptions []grpc.ServerOption
	a.grpcServer = grpc.NewServer(serverOptions...)

	protoCompanyService.RegisterCompanyServiceServer(a.grpcServer, server)
	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)

}
