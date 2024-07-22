package services

import (
	"testing"

	"github.com/safciplak/zippe-test-case/src/business/categories/models"
	"github.com/safciplak/zippe-test-case/src/business/categories/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService(t *testing.T) {
	repo := repositories.NewCategoryRepository()
	service := NewCategoryService(repo)

	t.Run("Create Category", func(t *testing.T) {
		category := models.Category{Name: "Test Category"}
		createdCategory := service.Create(category)

		assert.Equal(t, 1, createdCategory.ID)
		assert.Equal(t, "Test Category", createdCategory.Name)
	})

	t.Run("List Categories", func(t *testing.T) {
		categories := service.List()

		assert.Len(t, categories, 1)
		assert.Equal(t, 1, categories[0].ID)
		assert.Equal(t, "Test Category", categories[0].Name)
	})

	t.Run("Read Category", func(t *testing.T) {
		category, exists := service.Read(1)

		assert.True(t, exists)
		assert.Equal(t, 1, category.ID)
		assert.Equal(t, "Test Category", category.Name)
	})

	t.Run("Update Category", func(t *testing.T) {
		updatedCategory := models.Category{Name: "Updated Category"}
		category, exists := service.Update(1, updatedCategory)

		assert.True(t, exists)
		assert.Equal(t, 1, category.ID)
		assert.Equal(t, "Updated Category", category.Name)
	})

	t.Run("Delete Category", func(t *testing.T) {
		deleted := service.Delete(1)

		assert.True(t, deleted)

		categories := service.List()
		assert.Len(t, categories, 0)
	})

	t.Run("Read Non-existent Category", func(t *testing.T) {
		_, exists := service.Read(2)
		assert.False(t, exists)
	})

	t.Run("Update Non-existent Category", func(t *testing.T) {
		updatedCategory := models.Category{Name: "Non-existent Category"}
		_, exists := service.Update(2, updatedCategory)
		assert.False(t, exists)
	})

	t.Run("Delete Non-existent Category", func(t *testing.T) {
		deleted := service.Delete(2)
		assert.False(t, deleted)
	})
}
