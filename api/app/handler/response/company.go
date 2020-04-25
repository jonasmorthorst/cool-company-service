package response

import (
	"api/app/model"
	"sort"
	"time"
)

type SingleCompanyResponse struct {
	Company  *CompanyResponse `json:"company"`
}

type CompanyListResponse struct {
	Companies  []CompanyResponse `json:"companies"`
}

type CompanyResponse struct {
	Id   			uint  	`json:"id"`
	Name			string  `json:"name"`
	CompanyID		string  `json:"companyID"`
	Address			string  `json:"address"`
	Phone			string  `json:"phone"`
	Email			string  `json:"email"`
	City			string  `json:"city"`
	ZipCode			string  `json:"zipCode"`
	Country			string  `json:"country"`
	Owners			[]OwnerResponse  `json:"owners"`
	CreatedAt 		time.Time 	`json:"createdAt"`
}

type OwnerResponse struct {
	Id  	uint `json:"id"`
	Name  	string `json:"name"`
	Title  	string `json:"title"`
}

func NewCompanySingleResponse(m *model.Company) *SingleCompanyResponse {
	r := newCompanyResponse(m)

	return &SingleCompanyResponse{r}
}

func NewOwnersResponse(owners []*model.Owner) []OwnerResponse {
	or := OwnerResponse{}
	list := make([]OwnerResponse, 0)
	for _, o := range owners {
		or.Id = o.ID
		or.Name = o.Name
		or.Title = o.Title

		list = append(list, or)
	}

	return list
}

func newCompanyResponse(m *model.Company) *CompanyResponse {
	r := new(CompanyResponse)
	r.Id = m.ID
	r.Name = m.Name
	r.CompanyID = m.CompanyID
	r.Phone = m.Phone
	r.Email = m.Email
	r.Address = m.Address
	r.City = m.City
	r.ZipCode = m.ZipCode
	r.Country = m.Country
	r.Owners = NewOwnersResponse(m.Owners)

	r.CreatedAt = m.CreatedAt

	return r
}

func NewCompanyListResponse(venues []model.Company) *CompanyListResponse {
	r := new(CompanyListResponse)
	cr := CompanyResponse{}
	r.Companies = make([]CompanyResponse, 0)
	for _, i := range venues {
		cr.Id = i.ID
		cr.Name = i.Name
		cr.CompanyID = i.CompanyID
		cr.Phone = i.Phone
		cr.Email = i.Email
		cr.Address = i.Address
		cr.City = i.City
		cr.ZipCode = i.ZipCode
		cr.Country = i.Country
		cr.CreatedAt = i.CreatedAt

		r.Companies = append(r.Companies, cr)
	}

	//Sort by name
	sort.Slice(r.Companies, func(i, j int) bool {
		return r.Companies[i].Name < r.Companies[j].Name
	})

	return r
}