package service

import (
	"context"
	"mintapi/internal/core/domain"
)

//go:generate mockery --name Event

type Event interface {
	Create(ctx context.Context, event *domain.Event) (*domain.Event, error)
	Get(ctx context.Context, eventID uint64) (*domain.Event, error)
	Mint(ctx context.Context, ticketID uint64) (string, error)
}
