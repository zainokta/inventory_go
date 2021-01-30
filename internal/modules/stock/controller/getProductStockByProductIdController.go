package controller

import (
	"database/sql"
	"muramasa/internal/modules/stock/repository"
	"muramasa/internal/modules/stock/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProductStockByProductIdController struct {
	db *sql.DB
}

func NewGetProductStockByProductIdController(db *sql.DB) *GetProductStockByProductIdController {
	return &GetProductStockByProductIdController{db: db}
}

func (g *GetProductStockByProductIdController) GetProductStockByProductId(c *gin.Context) {
	stockRepository := repository.NewStockRepository(g.db)
	getProductStockByProductIdUseCase := usecase.NewGetProductStockByProductIdUseCase(stockRepository)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	productStock, err := getProductStockByProductIdUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    productStock,
	})
}
