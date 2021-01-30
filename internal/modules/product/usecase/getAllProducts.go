package usecase

import "muramasa/internal/modules/product/entity"

type GetAllProducts struct {
	productRepository entity.IProductRepository
}

func NewGetAllProductsUseCase(productRepository entity.IProductRepository) *GetAllProducts {
	return &GetAllProducts{productRepository: productRepository}
}

func (a *GetAllProducts) Execute() ([]*entity.Product, error) {
	products, err := a.productRepository.GetAllProduct()
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return []*entity.Product{}, nil
	}

	return products, nil
}
