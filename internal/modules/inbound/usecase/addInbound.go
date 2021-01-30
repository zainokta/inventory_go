package usecase

import "muramasa/internal/modules/inbound/entity"

type AddInbound struct {
	inboundRepository entity.IInboundRepository
}

func NewAddInbound(inboundRepository entity.IInboundRepository) AddInbound {
	return AddInbound{inboundRepository: inboundRepository}
}

func (a AddInbound) Execute(e entity.Inbound) {
	a.inboundRepository.AddInbound(e)
}
