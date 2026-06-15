package ledger

import (
	"context"

	"github.com/guisp/miniledger/internal/domain"
)

type TransferStore interface {
	FindByIdempotencyKey(ctx context.Context, key string) (*domain.Transfer, error)
	Save(ctx context.Context, t *domain.Transfer) error
}

type AccountStore interface {
	Get(ctx context.Context, id string) (*domain.Account, error)
}

type UnitOfWork interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}
