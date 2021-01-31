package controller

import (
	outbound "muramasa/internal/modules/outbound/entity"
	product "muramasa/internal/modules/product/entity"
	stock "muramasa/internal/modules/stock/entity"
)

type OrderController struct {
	stockRepository    stock.IStockRepository
	outboundRepository outbound.IOutboundRepository
	productRepository  product.IProductRepository
}

func NewOrderController(
	outboundRepository outbound.IOutboundRepository,
	stockRepository stock.IStockRepository,
	productRepository product.IProductRepository,
) *OrderController {
	return &OrderController{
		outboundRepository: outboundRepository,
		stockRepository:    stockRepository,
		productRepository:  productRepository,
	}
}
