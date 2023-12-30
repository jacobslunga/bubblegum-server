package service

import (
	"github.com/jacobslunga/bubblegum-server/internal/model"
	"github.com/jacobslunga/bubblegum-server/internal/repository"
)

type ComponentService struct {
	repo *repository.ComponentRepo
}

func NewComponentService(repo *repository.ComponentRepo) *ComponentService {
	return &ComponentService{repo: repo}
}

func (s *ComponentService) GetAllComponents() ([]model.Component, error) {
	return s.repo.FindAll()
}

func (s *ComponentService) GetComponentById(componentId string) (*model.Component, error) {
	return s.repo.FindById(componentId)
}

func (s *ComponentService) CreateComponent(component *model.Component) error {
	return s.repo.Create(component)
}

func (s *ComponentService) UpdateComponent(component *model.Component) error {
	return s.repo.Update(component)
}

func (s *ComponentService) DeleteComponent(component *model.Component) error {
	return s.repo.Delete(component)
}
