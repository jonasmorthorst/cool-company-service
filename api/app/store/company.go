package store

import (
	"github.com/jinzhu/gorm"
	"api/app/model"
)

type CompanyStore struct {
	db *gorm.DB
}

func NewCompanyStore(db *gorm.DB) *CompanyStore {
	return &CompanyStore{
		db: db,
	}
}

func (cs *CompanyStore) List() ([]model.Company, error) {
	var companies []model.Company

	err := cs.db.Find(&companies).Error

	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (vs *CompanyStore) GetByID(id uint) (*model.Company, error) {
	var m model.Company
	if err := vs.db.Preload("Owners").First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (vs *CompanyStore) Create(m *model.Company) (err error) {
	return vs.db.Create(m).Error
}

func (vs *CompanyStore) Update(m *model.Company) error {
	return vs.db.Model(m).Update(m).Error
}

func (vs *CompanyStore) Delete(m *model.Company) error {
	return vs.db.Model(m).Delete(m).Error
}