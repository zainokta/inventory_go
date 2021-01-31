package controller

import (
	"muramasa/internal/modules/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g *ProductController) GetAllProducts(c *gin.Context) {
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
