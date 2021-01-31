package usecase

import "muramasa/internal/modules/stock/entity"

type GetProductTotalStock struct {
	stockRepository entity.IStockRepository
}

func NewGetProductTotalStockUseCase(stockRepository entity.IStockRepository) *GetProductTotalStock {
	return &GetProductTotalStock{stockRepository: stockRepository}
}

func (g *GetProductTotalStock) Execute(productID int) (int, error) {
	stock, err := g.stockRepository.GetProductTotalStock(productID)
	if err != nil {
		return 0, err
	}

	return stock, nil
}
