package usecase

import "muramasa/internal/modules/outbound/entity"

type AddOutbound struct {
	outboundRepository entity.IOutboundRepository
}

func NewAddOutbound(outboundRepository entity.IOutboundRepository) *AddOutbound {
	return &AddOutbound{outboundRepository: outboundRepository}
}

func (a *AddOutbound) Execute(outbound *entity.Outbound) (int, error) {
	id, err := a.outboundRepository.InsertOutbound(outbound)

	if err != nil {
		return 0, err
	}

	return id, nil
}
