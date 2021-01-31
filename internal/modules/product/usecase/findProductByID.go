package usecase

import (
	"muramasa/internal/modules/product/entity"
)

type FindProductByIDUseCase struct {
	productRepository entity.IProductRepository
}

func NewFindProductByIdUseCase(productRepository entity.IProductRepository) *FindProductByIDUseCase {
	return &FindProductByIDUseCase{productRepository: productRepository}
}

func (f *FindProductByIDUseCase) Execute(id int) (*entity.Product, error) {
	product, err := f.productRepository.FindProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
