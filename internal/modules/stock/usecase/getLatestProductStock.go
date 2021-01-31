package usecase

import "muramasa/internal/modules/stock/entity"

type GetLatestProductStock struct {
	stockRepository entity.IStockRepository
}

func NewGetLatestProductStockUseCase(stockRepository entity.IStockRepository) *GetLatestProductStock {
	return &GetLatestProductStock{stockRepository: stockRepository}
}

func (g *GetLatestProductStock) Execute(productID int) (*entity.Stock, error) {
	stock, err := g.stockRepository.GetLatestProductStock(productID)
	if err != nil {
		return nil, err
	}

	return stock, nil
}
