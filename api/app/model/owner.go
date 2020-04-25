package model

import (
	"github.com/jinzhu/gorm"
)

type Owner struct {
	gorm.Model
	Name			string
	Title			string
	CompanyID		uint
	Company			Company
}
