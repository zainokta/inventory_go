package controller

import (
	inboundEntity "muramasa/internal/modules/inbound/entity"
	productEntity "muramasa/internal/modules/product/entity"
	stockEntity "muramasa/internal/modules/stock/entity"
)

type StockController struct {
	productRepository productEntity.IProductRepository
	inboundRepository inboundEntity.IInboundRepository
	stockRepository   stockEntity.IStockRepository
}

func NewStockController(
	productRepository productEntity.IProductRepository,
	inboundRepository inboundEntity.IInboundRepository,
	stockRepository stockEntity.IStockRepository,
) *StockController {
	return &StockController{
		productRepository: productRepository,
		inboundRepository: inboundRepository,
		stockRepository:   stockRepository,
	}
}
