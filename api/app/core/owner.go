package core

import "api/app/model"

type OwnerStore interface {
	GetByID(uint) (*model.Owner, error)
	Create(*model.Owner) error
	Update(*model.Owner) error
	Delete(*model.Owner) error
}
