package ticket

import (
	"context"
	"mintapi/internal/core/domain"
	"mintapi/internal/core/port/resource"
)

type Service struct {
	ContractRsc resource.Contract
}

func NewTicketService(contractRsc resource.Contract) *Service {
	return &Service{
		ContractRsc: contractRsc,
	}
}

func (s *Service) Create(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error) {
	return s.ContractRsc.CreateTicket(ctx, ticket)
}
