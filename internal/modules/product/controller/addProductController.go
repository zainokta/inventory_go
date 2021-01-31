package controller

import (
	"muramasa/internal/modules/product/entity"
	"muramasa/internal/modules/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddProductController struct {
	productRepository entity.IProductRepository
}

func NewAddProductController(productRepository entity.IProductRepository) *AddProductController {
	return &AddProductController{productRepository: productRepository}
}

func (g *AddProductController) AddProduct(c *gin.Context) {
	addProductUseCase := usecase.NewAddProductUseCase(g.productRepository)
	product := &entity.Product{}

	err := c.Bind(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	err = entity.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	id, err := addProductUseCase.Execute(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    id,
	})
}
