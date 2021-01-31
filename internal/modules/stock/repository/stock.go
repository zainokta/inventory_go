package repository

import (
	"database/sql"
	"muramasa/internal/modules/stock/entity"
)

type StockRepository struct {
	db *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (s *StockRepository) GetProductStockByProductId(id int) ([]*entity.Stock, error) {
	rows, err := s.db.Query(`
		SELECT stocks.*, products.*, inbounds.* 
		FROM stocks
		JOIN products ON products.id = stocks.product_id
		JOIN inbounds ON inbounds.id = stocks.inbound_id
		WHERE products.id = ?
	`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stocks []*entity.Stock
	hasResult := false
	for rows.Next() {
		hasResult = true
		stock := &entity.Stock{}
		err := rows.Scan(
			&stock.ID,
			&stock.ProductID,
			&stock.Stock,
			&stock.ExpiryDate,
			&stock.InboundID,
			&stock.Product.ID,
			&stock.Product.Name,
			&stock.Product.Sku,
			&stock.Product.Expirable,
			&stock.Inbound.ID,
			&stock.Inbound.InboundQuantity,
			&stock.Inbound.InboundDate,
			&stock.Inbound.Price,
			&stock.Inbound.TotalPrice,
			&stock.Inbound.InvoiceNo,
		)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	if !hasResult {
		return nil, nil
	}

	return stocks, nil
}

func (s *StockRepository) AddStock(stock *entity.Stock) (int, error) {
	stmt, err := s.db.Prepare("INSERT INTO stocks(product_id, stock, expiry_date, inbound_id) VALUES(?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(&stock.ProductID, &stock.Stock, &stock.ExpiryDate, &stock.InboundID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
