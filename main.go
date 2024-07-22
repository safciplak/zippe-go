package main

import (
	"github.com/gin-gonic/gin"
	categories "github.com/safciplak/zippe-test-case/src/business/categories/handlers"
	"github.com/safciplak/zippe-test-case/src/business/categories/repositories"
	"github.com/safciplak/zippe-test-case/src/business/categories/services"
	"log"
)

func main() {
	r := gin.Default()

	repo := repositories.NewCategoryRepository()
	svc := services.NewCategoryService(repo)
	categories := categories.NewCategoryHandler(svc)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/categories", categories.List)
		apiv1.GET("/categories/:id", categories.Read)
		apiv1.POST("/categories", categories.Create)
		apiv1.PUT("/categories/:id", categories.Update)
		apiv1.DELETE("/categories/:id", categories.Delete)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
