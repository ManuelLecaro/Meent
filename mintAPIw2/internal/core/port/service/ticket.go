package service

import (
	"context"
	"mintapi/internal/core/domain"
)

//go:generate mockery --name Ticket

type Ticket interface {
	Create(ctx context.Context, event *domain.Event) (*domain.Event, error)
}