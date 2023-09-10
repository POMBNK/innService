package service

import (
	"context"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/model"
	"github.com/POMBNK/shtrafovNetTestTask/pkg/client/rusprofile"
)

type CompanyService struct {
}

func NewService() *CompanyService {
	return &CompanyService{}
}

func (s *CompanyService) GetCompanyByInn(ctx context.Context, inn string) (model.Company, error) {
	rclient := rusprofile.NewClient()
	bytes, err := rclient.ParsePage(inn)
	if err != nil {
		return model.Company{}, err
	}

	info, err := rclient.ParseInfo(bytes)
	if err != nil {
		return model.Company{}, err
	}

	return model.NewCompany(
		info.Inn,
		info.Kpp,
		info.CompanyName,
		info.Fio), nil

}
