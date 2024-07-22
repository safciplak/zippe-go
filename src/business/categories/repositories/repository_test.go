package repositories

import (
	"testing"

	"github.com/safciplak/zippe-test-case/src/business/categories/models"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository(t *testing.T) {
	repo := NewCategoryRepository()

	t.Run("create category", func(t *testing.T) {
		category := models.Category{Name: "test category"}
		createdCategory := repo.Create(category)

		assert.Equal(t, 1, createdCategory.ID)
		assert.Equal(t, "test category", createdCategory.Name)
	})

	t.Run("list categories", func(t *testing.T) {
		categories := repo.List()

		assert.Len(t, categories, 1)
		assert.Equal(t, 1, categories[0].ID)
		assert.Equal(t, "test category", categories[0].Name)
	})

	t.Run("read category", func(t *testing.T) {
		category, exists := repo.Read(1)

		assert.True(t, exists)
		assert.Equal(t, 1, category.ID)
		assert.Equal(t, "test category", category.Name)
	})

	t.Run("update category", func(t *testing.T) {
		updatedCategory := models.Category{Name: "updated category"}
		category, exists := repo.Update(1, updatedCategory)

		assert.True(t, exists)
		assert.Equal(t, 1, category.ID)
		assert.Equal(t, "updated category", category.Name)
	})

	t.Run("delete category", func(t *testing.T) {
		deleted := repo.Delete(1)

		assert.True(t, deleted)

		categories := repo.List()
		assert.Len(t, categories, 0)
	})

	t.Run("read non-existent category", func(t *testing.T) {
		_, exists := repo.Read(2)
		assert.False(t, exists)
	})

	t.Run("update non-existent category", func(t *testing.T) {
		updatedCategory := models.Category{Name: "non-existent category"}
		_, exists := repo.Update(2, updatedCategory)
		assert.False(t, exists)
	})

	t.Run("delete non-existent category", func(t *testing.T) {
		deleted := repo.Delete(2)
		assert.False(t, deleted)
	})
}
