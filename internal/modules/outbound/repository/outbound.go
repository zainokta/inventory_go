package repository

import (
	"database/sql"
	"muramasa/internal/modules/outbound/entity"
)

type OutboundRepository struct {
	db *sql.DB
}

func NewOutboundRepository(db *sql.DB) *OutboundRepository {
	return &OutboundRepository{db: db}
}

func (o *OutboundRepository) InsertOutbound(outbound *entity.Outbound) (int, error) {
	stmt, err := o.db.Prepare("INSERT INTO outbounds(quantity, price, total_price, status, reference_number) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	totalPrice := outbound.Price * int64(outbound.Quantity)

	result, err := stmt.Exec(&outbound.Quantity, &outbound.Price, totalPrice, &outbound.Status, &outbound.ReferenceNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
