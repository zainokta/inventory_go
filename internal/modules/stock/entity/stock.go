package entity

import (
	inbound "muramasa/internal/modules/inbound/entity"
	product "muramasa/internal/modules/product/entity"
	"time"
)

type Stock struct {
	ID         int64           `json:"id"`
	ProductID  int64           `json:"product_id" binding:"required"`
	Stock      int             `json:"stock"`
	ExpiryDate *time.Time      `json:"expiry_date"`
	InboundID  int64           `json:"inbound_id"`
	Product    product.Product `json:"product"`
	Inbound    inbound.Inbound `json:"inbound"`
}

type IStockRepository interface {
	GetProductStockByProductId(int) ([]*Stock, error)
	AddStock(*Stock) (int, error)
	GetProductTotalStock(int) (int, error)
	GetLatestProductStock(int) (int, error)
	UpdateProductStock(int, int) error
}
