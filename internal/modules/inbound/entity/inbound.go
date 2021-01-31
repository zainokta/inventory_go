package entity

import "time"

type Inbound struct {
	ID              int64      `json:"id"`
	InboundQuantity int        `json:"inbound_quantity" binding:"required"`
	InboundDate     *time.Time `json:"inbound_date"`
	Price           int64      `json:"price" binding:"required"`
	TotalPrice      int64      `json:"total_price"`
	InvoiceNo       string     `json:"invoice_no" binding:"required"`
}

type IInboundRepository interface {
	AddInbound(*Inbound) (int, error)
}
