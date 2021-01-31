package controller

import (
	"fmt"
	outboundEntity "muramasa/internal/modules/outbound/entity"
	outboundUsecase "muramasa/internal/modules/outbound/usecase"
	productUseCase "muramasa/internal/modules/product/usecase"
	stockEntity "muramasa/internal/modules/stock/entity"
	stockUsecase "muramasa/internal/modules/stock/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderRequest struct {
	*outboundEntity.Outbound
	*stockEntity.Stock
}

func (o *OrderController) AddOrder(c *gin.Context) {
	findProductByIDUsecase := productUseCase.NewFindProductByIdUseCase(o.productRepository)
	getProductTotalStockUseCase := stockUsecase.NewGetProductTotalStockUseCase(o.stockRepository)
	getLatestProductStockUseCase := stockUsecase.NewGetLatestProductStockUseCase(o.stockRepository)
	updateProductStockUseCase := stockUsecase.NewUpdateProductStockUseCase(o.stockRepository)
	addOutboundUseCase := outboundUsecase.NewAddOutbound(o.outboundRepository)

	request := &orderRequest{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	product, err := findProductByIDUsecase.Execute(int(request.Stock.ProductID))
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

	totalStock, err := getProductTotalStockUseCase.Execute(int(request.Stock.ProductID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	requestedQuantity := request.Outbound.Quantity

	if totalStock < requestedQuantity {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    "Ordered quantity cannot be more than total product stock.",
		})
		return
	}

	for stock, err := getLatestProductStockUseCase.Execute(int(request.Stock.ProductID)); requestedQuantity != 0; stock, err = getLatestProductStockUseCase.Execute(int(request.Stock.ProductID)) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err.Error(),
			})
			return
		}

		currentStock := stock.Stock
		if currentStock < requestedQuantity {
			requestedQuantity -= currentStock
			currentStock = 0
		} else {
			currentStock -= requestedQuantity
			requestedQuantity = 0
		}

		err = updateProductStockUseCase.Execute(int(stock.ID), currentStock)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
				"data":    err.Error(),
			})
			return
		}
	}

	_, err = addOutboundUseCase.Execute(request.Outbound)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    "Success creating outbound for order",
	})
}
