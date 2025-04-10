package shared_ports

import (
	"context"

	shared_domain "github.com/DongnutLa/stockio/internal/zshared/core/domain"
)

type IConsecutiveService interface {
	GetConsecutive(ctx context.Context, tipo shared_domain.ConsecutiveType) (string, error)
	RestoreConsecutive(ctx context.Context, tipo shared_domain.ConsecutiveType) error
}
