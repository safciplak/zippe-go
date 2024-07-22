package services

import (
	"github.com/safciplak/zippe-test-case/src/business/categories/models"
	"github.com/safciplak/zippe-test-case/src/business/categories/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List() []models.Category {
	return s.repo.List()
}

func (s *CategoryService) Read(id int) (models.Category, bool) {
	return s.repo.Read(id)
}

func (s *CategoryService) Create(category models.Category) models.Category {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(id int, category models.Category) (models.Category, bool) {
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int) bool {
	return s.repo.Delete(id)
}
