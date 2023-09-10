package company

import (
	"context"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/model"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/service"
)

type Policy struct {
	companyService *service.CompanyService
}

func NewPolicy(companyService *service.CompanyService) *Policy {
	return &Policy{
		companyService: companyService,
	}
}

func (p *Policy) GetCompanyByInn(ctx context.Context, inn string) (model.Company, error) {

	company, err := p.companyService.GetCompanyByInn(ctx, inn)
	if err != nil {
		return model.Company{}, err
	}

	return company, nil
}
