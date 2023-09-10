package app

import (
	"context"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	grpcCompanyServer "github.com/POMBNK/shtrafovNetTestTask/internal/controller/grpc/v1/company"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/service"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/policy/company"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net/http"
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

}

func (a *App) startGRPC(ctx context.Context, server interface{}) error {

}
