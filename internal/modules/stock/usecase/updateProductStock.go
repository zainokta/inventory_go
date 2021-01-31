package usecase

import "muramasa/internal/modules/stock/entity"

type UpdateProductStock struct {
	stockRepository entity.IStockRepository
}

func NewUpdateProductStockUseCase(stockRepository entity.IStockRepository) *UpdateProductStock {
	return &UpdateProductStock{stockRepository: stockRepository}
}

func (g *UpdateProductStock) Execute(stockID int, currentStock int) error {
	err := g.stockRepository.UpdateProductStock(stockID, currentStock)
	if err != nil {
		return err
	}

	return nil
}
