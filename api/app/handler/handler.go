package handler

import (
	"api/app/core"
)

type Handler struct {
	companyStore      	core.CompanyStore
	ownerStore       	core.OwnerStore
}

func NewHandler(cs core.CompanyStore, os core.OwnerStore) *Handler {
	return &Handler{
		companyStore: cs,
		ownerStore:  os,
	}
}