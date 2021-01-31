package controller

import (
	"muramasa/internal/modules/product/entity"
	"muramasa/internal/modules/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	productRepository entity.IProductRepository
}

func NewGetAllProductsController(productRepository entity.IProductRepository) *GetAllProductsController {
	return &GetAllProductsController{productRepository: productRepository}
}

func (g *GetAllProductsController) GetAllProducts(c *gin.Context) {
	getAllProductUseCase := usecase.NewGetAllProductsUseCase(g.productRepository)
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
