package resource

import (
	"context"
	"mintapi/internal/core/domain"
)

//go:generate mockery --name Contract

type Contract interface {
	CreateEvent(ctx context.Context, event *domain.Event) (*domain.Event, error)
	GetEvent(ctx context.Context, eventID uint64) (*domain.Event, error)
	
}
