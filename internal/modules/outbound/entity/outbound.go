package entity

import "time"

type Outbound struct {
	ID              int64      `json:"id"`
	Quantity        int        `json:"quantity"`
	Price           int64      `json:"price"`
	TotalPrice      int64      `json:"total_price"`
	ReferenceNumber string     `json:"reference_number"`
	Timestamp       *time.Time `json:"timestamp"`
}

type IOutboundRepository interface {
	AddOutbound(Outbound) int
}
