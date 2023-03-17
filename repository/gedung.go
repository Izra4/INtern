package repository

import (
	"InternBCC/entity"
	"gorm.io/gorm"
)

type GedungRepository interface {
	FindAll() ([]entity.Gedung, error)
	FindByID(id uint) (entity.Gedung, error)
}

type gedungRepository struct {
	db *gorm.DB
}

func NewGedungRepository(db *gorm.DB) GedungRepository {
	return &gedungRepository{db: db}
}

func (r *gedungRepository) FindAll() ([]entity.Gedung, error) {
	var get []entity.Gedung
	if err := r.db.Debug().Preload("Links").Find(&get).Error; err != nil {
		return nil, err
	}
	return get, nil
}

func (r *gedungRepository) FindByID(id uint) (entity.Gedung, error) {
	var get entity.Gedung
	if err := r.db.Preload("Links").First(&get, id).Error; err != nil {
		return entity.Gedung{}, err
	}
	return get, nil
}
