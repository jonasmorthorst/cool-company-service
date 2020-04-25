package store

import (
	"github.com/jinzhu/gorm"
	"api/app/model"
)

type OwnerStore struct {
	db *gorm.DB
}

func NewOwnerStore(db *gorm.DB) *OwnerStore {
	return &OwnerStore{
		db: db,
	}
}

func (us *OwnerStore) GetByID(id uint) (*model.Owner, error) {
	var m model.Owner
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *OwnerStore) Create(o *model.Owner) (err error) {
	return us.db.Create(o).Error
}

func (us *OwnerStore) Update(o *model.Owner) error {
	return us.db.Model(o).Update(o).Error
}

func (vs *OwnerStore) Delete(o *model.Owner) error {
	return vs.db.Model(o).Delete(o).Error
}
