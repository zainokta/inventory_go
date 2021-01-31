package repository

import (
	"database/sql"
	"muramasa/internal/modules/product/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) GetAllProduct() ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*entity.Product{}

	for rows.Next() {
		product := &entity.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Sku, &product.Expirable)
		if err != nil {
			return nil, err
		}

		result = append(result, product)
	}

	return result, nil
}

func (p ProductRepository) FindProductByID(id int) (*entity.Product, error) {
	rows, err := p.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	product := &entity.Product{}
	hasResult := false
	for rows.Next() {
		hasResult = true
		err := rows.Scan(&product.ID, &product.Name, &product.Sku, &product.Expirable)
		if err != nil {
			return nil, err
		}
	}

	if !hasResult {
		return nil, nil
	}

	return product, nil

}

func (p ProductRepository) AddProduct(product *entity.Product) (int, error) {
	stmt, err := p.db.Prepare("INSERT INTO products(name,sku,expireable) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(&product.Name, &product.Sku, &product.Expirable)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
