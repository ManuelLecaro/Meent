package event

import (
	"context"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/resource"
)

type Service struct {
	ContractRsc resource.Contract
}

func NewEventService(contractRsc resource.Contract) *Service {
	return &Service{
		ContractRsc: contractRsc,
	}
}

func (s *Service) Create(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	return s.ContractRsc.CreateEvent(ctx, event)
}
