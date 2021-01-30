package usecase

import "muramasa/internal/modules/product/entity"

type AddProduct struct {
	productRepository entity.IProductRepository
}

func NewAddProductUseCase(productRepository entity.IProductRepository) *AddProduct {
	return &AddProduct{productRepository: productRepository}
}

func (a *AddProduct) Execute(product *entity.Product) (int, error) {
	id, err := a.productRepository.AddProduct(product)

	if err != nil {
		return 0, err
	}

	return id, nil
}
