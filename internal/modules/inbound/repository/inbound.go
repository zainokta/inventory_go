package repository

import "database/sql"

type InboundRepository struct {
	db *sql.DB
}

func NewInboundRepository(db *sql.DB) InboundRepository {
	return InboundRepository{db: db}
}

func AddInbound() int {
	return 1
}
