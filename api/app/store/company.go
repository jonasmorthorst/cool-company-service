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

func (cs *CompanyStore) GetByID(id uint) (*model.Company, error) {
	var c model.Company
	if err := cs.db.Preload("Owners").First(&c, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil
}

func (cs *CompanyStore) Create(c *model.Company) (err error) {
	return cs.db.Create(c).Error
}

func (cs *CompanyStore) Update(c *model.Company) error {
	return cs.db.Model(c).Update(c).Error
}

func (cs *CompanyStore) Delete(c *model.Company) error {
	// First delete all owners
	err := cs.db.Delete(model.Owner{}, "company_id = ?", c.ID).Error

	if err != nil {
		return nil
	}

	return cs.db.Model(c).Delete(c).Error
}