package model

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name			string
	CompanyID		string
	Email			string
	Phone			string
	Address			string
	ZipCode			string
	City			string
	Country			string
	Owners 			[]*Owner
}