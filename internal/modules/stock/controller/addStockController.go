package controller

import (
	"fmt"
	inboundEntity "muramasa/internal/modules/inbound/entity"
	inboundUsecase "muramasa/internal/modules/inbound/usecase"
	productUseCase "muramasa/internal/modules/product/usecase"
	stockEntity "muramasa/internal/modules/stock/entity"
	stockUsecase "muramasa/internal/modules/stock/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stockRequest struct {
	*stockEntity.Stock
	*inboundEntity.Inbound
}

func (s *StockController) AddStock(c *gin.Context) {
	productUseCase := productUseCase.NewFindProductByIdUseCase(s.productRepository)
	addInboundUseCase := inboundUsecase.NewAddInboundUseCase(s.inboundRepository)
	addStockUseCase := stockUsecase.NewAddStockUseCase(s.stockRepository)

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
