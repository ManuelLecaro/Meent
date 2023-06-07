package resource

import (
	"context"
	"mintapi/internal/core/domain"
)

//go:generate mockery --name Minter
type Minter interface {
	UploadToIPFS(ctx context.Context, data *domain.Event) (string, error)
}
