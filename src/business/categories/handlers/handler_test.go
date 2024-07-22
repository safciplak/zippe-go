package categories

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/safciplak/zippe-test-case/src/business/categories/models"
	"github.com/safciplak/zippe-test-case/src/business/categories/repositories"
	"github.com/safciplak/zippe-test-case/src/business/categories/services"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	repo := repositories.NewCategoryRepository()
	service := services.NewCategoryService(repo)
	handler := NewCategoryHandler(service)

	router := gin.Default()
	router.GET("/categories", handler.List)
	router.GET("/categories/:id", handler.Read)
	router.POST("/categories", handler.Create)
	router.PUT("/categories/:id", handler.Update)
	router.DELETE("/categories/:id", handler.Delete)

	return router
}

func TestCategoryHandler(t *testing.T) {
	router := setupRouter()

	t.Run("list categories", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/categories", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("create category", func(t *testing.T) {
		category := models.Category{Name: "test category"}
		data, _ := json.Marshal(category)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var createdCategory models.Category
		json.Unmarshal(w.Body.Bytes(), &createdCategory)
		assert.Equal(t, "test category", createdCategory.Name)
	})

	t.Run("read category", func(t *testing.T) {
		category := models.Category{Name: "read test category"}
		data, _ := json.Marshal(category)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		var createdCategory models.Category
		json.Unmarshal(w.Body.Bytes(), &createdCategory)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodGet, "/categories/"+strconv.Itoa(createdCategory.ID), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var readCategory models.Category
		json.Unmarshal(w.Body.Bytes(), &readCategory)
		assert.Equal(t, createdCategory.ID, readCategory.ID)
		assert.Equal(t, "read test category", readCategory.Name)
	})

	t.Run("update category", func(t *testing.T) {
		category := models.Category{Name: "update test category"}
		data, _ := json.Marshal(category)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		var createdCategory models.Category
		json.Unmarshal(w.Body.Bytes(), &createdCategory)

		updatedCategory := models.Category{Name: "updated category"}
		data, _ = json.Marshal(updatedCategory)
		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodPut, "/categories/"+strconv.Itoa(createdCategory.ID), bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var updatedCategoryResponse models.Category
		json.Unmarshal(w.Body.Bytes(), &updatedCategoryResponse)
		assert.Equal(t, "updated category", updatedCategoryResponse.Name)
	})

	t.Run("delete category", func(t *testing.T) {
		category := models.Category{Name: "delete test category"}
		data, _ := json.Marshal(category)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/categories", bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		var createdCategory models.Category
		json.Unmarshal(w.Body.Bytes(), &createdCategory)

		w = httptest.NewRecorder()
		categoryId := strconv.Itoa(createdCategory.ID)
		req, _ = http.NewRequest(http.MethodDelete, "/categories/"+categoryId, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest(http.MethodGet, "/categories/"+categoryId, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
