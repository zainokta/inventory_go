package infrastructure

import (
	"database/sql"
	inbound "muramasa/internal/modules/inbound/repository"
	outbound "muramasa/internal/modules/outbound/repository"
	product "muramasa/internal/modules/product/repository"
	stock "muramasa/internal/modules/stock/repository"

	inboundEntity "muramasa/internal/modules/inbound/entity"
	outboundEntity "muramasa/internal/modules/outbound/entity"
	productEntity "muramasa/internal/modules/product/entity"
	stockEntity "muramasa/internal/modules/stock/entity"
)

type Repositories struct {
	productRepository  productEntity.IProductRepository
	stockRepository    stockEntity.IStockRepository
	inboundRepository  inboundEntity.IInboundRepository
	outboundRepository outboundEntity.IOutboundRepository
}

func initRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		productRepository:  product.NewProductRepository(db),
		stockRepository:    stock.NewStockRepository(db),
		inboundRepository:  inbound.NewInboundRepository(db),
		outboundRepository: outbound.NewOutboundRepository(db),
	}
}
