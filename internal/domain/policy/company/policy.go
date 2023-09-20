package company

import (
	"context"
	"errors"
	"github.com/POMBNK/shtrafovNetTestTask/internal/apperror"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/model"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/service"
	"net/http"
	"strconv"
)

var ErrInnInvalid = apperror.NewAppError(http.StatusBadRequest, "invalid inn")

type Policy struct {
	companyService *service.CompanyService
}

func NewPolicy(companyService *service.CompanyService) *Policy {
	return &Policy{
		companyService: companyService,
	}
}

func (p *Policy) GetCompanyByInn(ctx context.Context, inn string) (model.Company, error) {
	innRunes := []rune(inn)
	err := validateInn(string(innRunes))
	if err != nil {
		if errors.Is(err, ErrInnInvalid) {
			return model.Company{}, err
		}
		return model.Company{}, err
	}

	company, err := p.companyService.GetCompanyByInn(ctx, inn)
	if err != nil {
		return model.Company{}, err
	}

	return company, nil
}

func validateInn(inn string) error {
	if len(inn) < 10 || len(inn) > 12 {
		return ErrInnInvalid
	}

	if _, err := strconv.Atoi(inn); err != nil {
		return ErrInnInvalid
	}

	return nil
}
