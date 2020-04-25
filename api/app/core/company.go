package core

import "api/app/model"

type CompanyStore interface {
	List() ([]model.Company, error)
	GetByID(uint) (*model.Company, error)
	Create(*model.Company) error
	Update(*model.Company) error
	Delete(*model.Company) error
}
