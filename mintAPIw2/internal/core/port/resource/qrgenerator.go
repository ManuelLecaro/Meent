package resource

import "context"

//go:generate mockery --name QRGenerator

type QRGenerator interface {
	CreateQR(ctx context.Context, data map[string]interface{}, name string) error
}
