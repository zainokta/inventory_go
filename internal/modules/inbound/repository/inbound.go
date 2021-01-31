package repository

import (
	"database/sql"
	"muramasa/internal/modules/inbound/entity"
)

type InboundRepository struct {
	db *sql.DB
}

func NewInboundRepository(db *sql.DB) *InboundRepository {
	return &InboundRepository{db: db}
}

func (i *InboundRepository) AddInbound(inbound *entity.Inbound) (int, error) {
	stmt, err := i.db.Prepare("INSERT INTO inbounds(inbound_quantity, price, total_price, invoice_no) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	totalPrice := int64(inbound.InboundQuantity) * inbound.Price

	result, err := stmt.Exec(&inbound.InboundQuantity, &inbound.Price, totalPrice, &inbound.InvoiceNo)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
