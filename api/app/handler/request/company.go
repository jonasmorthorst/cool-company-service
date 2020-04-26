package request

import (
	"api/app/model"
	"github.com/labstack/echo/v4"
)

type CompanyRequest struct {
	Name			string 	`json:"name" validate:"required"`
	CompanyID		string	`json:"companyID" validate:"required"`
	Email			string	`json:"email"`
	Phone			string	`json:"phone"`
	Address			string	`json:"address" validate:"required"`
	ZipCode			string	`json:"zipCode" validate:"required"`
	City			string	`json:"city" validate:"required"`
	Country			string	`json:"country" validate:"required"`
}

type OwnerRequest struct {
	Name  	string `json:"name" validate:"required"`
	Title  	string `json:"title" validate:"required"`
}

func (r *CompanyRequest) Populate(c *model.Company) {
	r.Name = c.Name
	r.CompanyID = r.CompanyID
	r.Email = c.Email
	r.Phone = c.Phone
	r.Address = c.Address
	r.ZipCode = c.ZipCode
	r.City = c.City
	r.Country = c.Country
}

func (r *CompanyRequest) Bind(c echo.Context, m *model.Company) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	m.Name = r.Name
	m.CompanyID = r.CompanyID
	m.Email = r.Email
	m.Phone = r.Phone
	m.Address = r.Address
	m.ZipCode = r.ZipCode
	m.City = r.City
	m.Country = r.Country

	return nil
}

func (r *OwnerRequest) Bind(c echo.Context, o *model.Owner) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	o.Name = r.Name
	o.Title = r.Title

	return nil
}