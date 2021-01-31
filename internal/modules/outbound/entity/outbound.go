package entity

import "time"

type Outbound struct {
	ID              int64      `json:"id"`
	Quantity        int        `json:"quantity" binding:"required"`
	Price           int64      `json:"price" binding:"required"`
	TotalPrice      int64      `json:"total_price"`
	Status          string     `json:"status" binding:"required"`
	ReferenceNumber string     `json:"reference_number" binding:"required"`
	Timestamp       *time.Time `json:"timestamp"`
}

type IOutboundRepository interface {
	InsertOutbound(*Outbound) (int, error)
}
