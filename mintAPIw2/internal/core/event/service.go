package event

import (
	"context"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/resource"
	"mintapi/internal/core/port/service"
)

type Service struct {
	ContractRsc resource.Contract
	Minter      resource.Minter
}

func NewEventService(contractRsc resource.Contract, minter resource.Minter) service.Event {
	return &Service{
		ContractRsc: contractRsc,
		Minter:      minter,
	}
}

func (s *Service) Create(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	return s.ContractRsc.CreateEvent(ctx, event)
}

func (s *Service) Get(ctx context.Context, eventID uint64) (*domain.Event, error) {
	return s.ContractRsc.GetEvent(ctx, eventID)
}

func (s *Service) Mint(ctx context.Context, eventID uint64) (string, error) {
	event, err := s.ContractRsc.GetEvent(ctx, eventID)
	if err != nil {
		return "", err
	}

	return s.Minter.UploadToIPFS(ctx, event)
}
