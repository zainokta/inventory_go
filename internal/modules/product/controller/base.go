package controller

import "muramasa/internal/modules/product/entity"

type ProductController struct {
	productRepository entity.IProductRepository
}

func NewProductController(productRepository entity.IProductRepository) *ProductController {
	return &ProductController{productRepository: productRepository}
}
