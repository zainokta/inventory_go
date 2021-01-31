package infrastructure

import (
	"database/sql"
	inbound "muramasa/internal/modules/inbound/repository"
	product "muramasa/internal/modules/product/repository"
	stock "muramasa/internal/modules/stock/repository"

	inboundEntity "muramasa/internal/modules/inbound/entity"
	productEntity "muramasa/internal/modules/product/entity"
	stockEntity "muramasa/internal/modules/stock/entity"
)

type Repositories struct {
	productRepository productEntity.IProductRepository
	stockRepository   stockEntity.IStockRepository
	inboundRepository inboundEntity.IInboundRepository
}

func initRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		productRepository: product.NewProductRepository(db),
		stockRepository:   stock.NewStockRepository(db),
		inboundRepository: inbound.NewInboundRepository(db),
	}
}
