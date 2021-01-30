package usecase

import "muramasa/internal/modules/stock/entity"

type GetProductStockByProductId struct {
	stockRepository entity.IStockRepository
}

func NewGetProductStockByProductIdUseCase(stockRepository entity.IStockRepository) *GetProductStockByProductId {
	return &GetProductStockByProductId{stockRepository: stockRepository}
}

func (g *GetProductStockByProductId) Execute(id int) ([]*entity.Stock, error) {
	stocks, err := g.stockRepository.GetProductStockByProductId(id)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
