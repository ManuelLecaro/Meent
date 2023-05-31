package service

import (
	"context"
	"mintapi/internal/core/domain"
)

//go:generate mockery --name Ticket

type Ticket interface {
	Create(ctx context.Context, ticket *domain.Ticket) (*domain.Ticket, error)
}
