package company

import (
	protoCompanyModel "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/model/v1"
	protoCompanyService "github.com/POMBNK/shtrafovNetTestTask/gen/go/inn_service/service/v1"
	"github.com/POMBNK/shtrafovNetTestTask/internal/domain/company/model"
)

func NewGetCompanyResponse(in model.Company) *protoCompanyService.GetCompanyResponse {
	return &protoCompanyService.GetCompanyResponse{
		Company: NewCompanyPB(in),
	}
}

func NewCompanyPB(entity model.Company) *protoCompanyModel.Company {
	pbCompany := &protoCompanyModel.Company{
		Inn:         entity.Inn,
		Kpp:         entity.Kpp,
		CompanyName: entity.CompanyName,
		Fio:         entity.Fio,
	}

	return pbCompany
}
