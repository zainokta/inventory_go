package usecase

import (
	"muramasa/internal/modules/inbound/entity"
)

type AddInboundUseCase struct {
	inboundRepository entity.IInboundRepository
}

func NewAddInboundUseCase(inboundRepository entity.IInboundRepository) *AddInboundUseCase {
	return &AddInboundUseCase{inboundRepository: inboundRepository}
}

func (a *AddInboundUseCase) Execute(inbound *entity.Inbound) (int, error) {
	id, err := a.inboundRepository.AddInbound(inbound)
	if err != nil {
		return 0, err
	}

	return id, nil
}
