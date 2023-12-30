package repository

import (
	"github.com/jacobslunga/bubblegum-server/internal/model"
	"gorm.io/gorm"
)

type ComponentRepo struct {
	DB *gorm.DB
}

func NewComponentRepo(db *gorm.DB) *ComponentRepo {
	return &ComponentRepo{
		DB: db,
	}
}

func (repo *ComponentRepo) FindAll() ([]model.Component, error) {
	var components []model.Component

	result := repo.DB.Find(&components)

	if result.Error != nil {
		return nil, result.Error
	}

	return components, nil
}

func (repo *ComponentRepo) FindById(id string) (*model.Component, error) {
	var component model.Component

	result := repo.DB.First(&component, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &component, nil
}

func (repo *ComponentRepo) Create(component *model.Component) error {
	result := repo.DB.Create(component)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ComponentRepo) Delete(component *model.Component) error {
	result := repo.DB.Delete(component)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ComponentRepo) Update(component *model.Component) error {
	result := repo.DB.Save(component)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
