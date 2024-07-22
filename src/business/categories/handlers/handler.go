package categories

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/safciplak/zippe-test-case/src/business/categories/models"
	"github.com/safciplak/zippe-test-case/src/business/categories/services"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) List(c *gin.Context) {
	log.Println("listing all categories")
	catList := h.service.List()
	c.JSON(http.StatusOK, catList)
}

func (h *CategoryHandler) Read(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("error reading category ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	log.Printf("reading category with ID: %d\n", id)
	category, exists := h.service.Read(id)
	if !exists {
		log.Printf("category with ID %d not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		log.Printf("error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	createdCategory := h.service.Create(newCategory)
	log.Printf("created new category with ID: %d\n", createdCategory.ID)
	c.JSON(http.StatusCreated, createdCategory)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("error reading category ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	var updatedCategory models.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		log.Printf("error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	updatedCategory, exists := h.service.Update(id, updatedCategory)
	if !exists {
		log.Printf("category with ID %d not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	log.Printf("updated category with ID: %d\n", id)
	c.JSON(http.StatusOK, updatedCategory)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("error reading category ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category id"})
		return
	}

	if !h.service.Delete(id) {
		log.Printf("category with ID %d not found\n", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	log.Printf("deleted category with ID: %d\n", id)
	c.Status(http.StatusNoContent)
}
