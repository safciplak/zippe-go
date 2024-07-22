package repositories

import (
	"sync"

	"github.com/safciplak/zippe-test-case/src/business/categories/models"
)

type CategoryRepository struct {
	categories map[int]models.Category
	nextID     int
	mu         sync.Mutex
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: make(map[int]models.Category),
		nextID:     1,
	}
}

func (r *CategoryRepository) List() []models.Category {
	r.mu.Lock()
	defer r.mu.Unlock()

	var catList []models.Category
	for _, category := range r.categories {
		catList = append(catList, category)
	}
	return catList
}

func (r *CategoryRepository) Read(id int) (models.Category, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	category, exists := r.categories[id]
	return category, exists
}

func (r *CategoryRepository) Create(category models.Category) models.Category {
	r.mu.Lock()
	defer r.mu.Unlock()

	category.ID = r.nextID
	r.nextID++
	r.categories[category.ID] = category
	return category
}

func (r *CategoryRepository) Update(id int, category models.Category) (models.Category, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[id]; !exists {
		return models.Category{}, false
	}

	category.ID = id
	r.categories[id] = category
	return category, true
}

func (r *CategoryRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[id]; !exists {
		return false
	}

	delete(r.categories, id)
	return true
}
