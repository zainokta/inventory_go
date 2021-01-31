package controller

import (
	"database/sql"
	"fmt"
	inboundEntity "muramasa/internal/modules/inbound/entity"
	inbound "muramasa/internal/modules/inbound/repository"
	inboundUsecase "muramasa/internal/modules/inbound/usecase"
	product "muramasa/internal/modules/product/repository"
	productUseCase "muramasa/internal/modules/product/usecase"
	stockEntity "muramasa/internal/modules/stock/entity"
	stock "muramasa/internal/modules/stock/repository"
	stockUsecase "muramasa/internal/modules/stock/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stockRequest struct {
	*stockEntity.Stock
	*inboundEntity.Inbound
}

type AddStockController struct {
	db *sql.DB
}

func NewAddStockController(db *sql.DB) *AddStockController {
	return &AddStockController{db: db}
}

func (a *AddStockController) AddStock(c *gin.Context) {
	productRepository := product.NewProductRepository(a.db)
	inboundRepository := inbound.NewInboundRepository(a.db)
	stockRepository := stock.NewStockRepository(a.db)

	productUseCase := productUseCase.NewFindProductByIdUseCase(productRepository)
	addInboundUseCase := inboundUsecase.NewAddInboundUseCase(inboundRepository)
	addStockUseCase := stockUsecase.NewAddStockUseCase(stockRepository)

	request := &stockRequest{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	product, err := productUseCase.Execute(int(request.Stock.ProductID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"data":    fmt.Sprintf("Product with id = %d not found", request.Stock.ProductID),
		})
		return
	}

	inboundID, err := addInboundUseCase.Execute(request.Inbound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	request.Stock.Stock = request.InboundQuantity
	request.Stock.InboundID = int64(inboundID)

	stockID, err := addStockUseCase.Execute(request.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    stockID,
	})
}
