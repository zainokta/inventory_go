package usecase

import (
	"muramasa/internal/modules/stock/entity"
	stockEntity "muramasa/internal/modules/stock/entity"
)

type AddStockUseCase struct {
	stockRepository stockEntity.IStockRepository
}

func NewAddStockUseCase(stockRepository stockEntity.IStockRepository) *AddStockUseCase {
	return &AddStockUseCase{
		stockRepository: stockRepository,
	}
}

func (s *AddStockUseCase) Execute(stock *entity.Stock) (int, error) {
	id, err := s.stockRepository.AddStock(stock)
	if err != nil {
		return 0, err
	}

	return id, nil
}
