package model

type Company struct {
	Inn         string
	Kpp         string
	CompanyName string
	Fio         string
}

func NewCompany(inn, kpp, companyName, fio string) Company {
	return Company{
		Inn:         inn,
		Kpp:         kpp,
		CompanyName: companyName,
		Fio:         fio,
	}
}
