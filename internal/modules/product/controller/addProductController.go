package controller

import (
	"database/sql"
	"muramasa/internal/modules/product/entity"
	"muramasa/internal/modules/product/repository"
	"muramasa/internal/modules/product/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddProductController struct {
	DB *sql.DB
}

func NewAddProductController(db *sql.DB) *AddProductController {
	return &AddProductController{DB: db}
}

func (g *AddProductController) AddProduct(c *gin.Context) {
	productRepository := repository.NewProductRepository(g.DB)
	addProductUseCase := usecase.NewAddProductUseCase(productRepository)
	product := &entity.Product{}

	err := c.Bind(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	err = entity.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := addProductUseCase.Execute(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product has been added.",
		"data":    id,
	})
}
