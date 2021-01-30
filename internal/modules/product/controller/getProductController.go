package controller

import (
	"database/sql"
	"muramasa/internal/modules/product/repository"
	"muramasa/internal/modules/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	DB *sql.DB
}

func NewGetAllProductsController(db *sql.DB) *GetAllProductsController {
	return &GetAllProductsController{DB: db}
}

func (g *GetAllProductsController) GetAllProducts(c *gin.Context) {
	productRepository := repository.NewProductRepository(g.DB)
	getAllProductUseCase := usecase.NewGetAllProductsUseCase(productRepository)
	products, err := getAllProductUseCase.Execute()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    products,
	})
}
