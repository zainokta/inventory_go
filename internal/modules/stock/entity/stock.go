package stock

import (
	inbound "muramasa/internal/modules/inbound/entity"
	product "muramasa/internal/modules/product/entity"
	"time"
)

type Stock struct {
	ID         int64           `json:"id"`
	ProductID  int64           `json:"product_id"`
	Product    product.Product `json:"product"`
	Stock      int             `json:"stock"`
	ExpiryDate *time.Time      `json:"expiry_date"`
	InboundID  int64           `json:"inbound_id"`
	Inbound    inbound.Inbound `json:"inbound"`
}

type IStockRepository interface {
	AddStock(Stock) int
}
