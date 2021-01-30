package entity

import "time"

type Inbound struct {
	ID              int64      `json:"id"`
	InboundQuantity int        `json:"inbound_quantity"`
	InboundDate     *time.Time `json:"inbound_date"`
	Price           int64      `json:"price"`
	TotalPrice      int64      `json:"total_price"`
	InvoiceNo       string     `json:"invoice_no"`
}

type IInboundRepository interface {
	AddInbound(Inbound) int
}