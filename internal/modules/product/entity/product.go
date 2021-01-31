package entity

import (
	"errors"
	"reflect"
)

type Product struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" binding:"required"`
	Sku       string `json:"sku" binding:"required"`
	Expirable bool   `json:"expirable" binding:"required"`
}

type IProductRepository interface {
	GetAllProduct() ([]*Product, error)
	AddProduct(*Product) (int, error)
	FindProductByID(int) (*Product, error)
}

func CreateProduct(product *Product) error {
	err := name(product.Name)
	if err != nil {
		return err
	}

	err = sku(product.Sku)
	if err != nil {
		return err
	}

	err = expireable(product.Expirable)
	if err != nil {
		return err
	}

	return nil
}

func name(name string) error {
	if name == "" {
		return errors.New("Product name cannot be empty.")
	}

	return nil
}

func sku(sku string) error {
	if sku == "" {
		return errors.New("Product SKU cannot be empty")
	}

	return nil
}

func expireable(expireable bool) error {
	if reflect.TypeOf(expireable).Kind() != reflect.Bool {
		return errors.New("Product Expireable must be a boolean type")
	}

	return nil
}
