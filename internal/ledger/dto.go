package ledger

import (
	"context"
	"fmt"

	"github.com/guisp/miniledger/internal/domain"
)

type ProcessTransferInput struct {
	From, To       domain.AccountID
	Money          domain.Money
	IdempotencyKey string
}

func (l *Ledger) ProcessTransfer(ctx context.Context, in ProcessTransferInput) (*domain.Transfer, error) {
	// 1. Idempotency: have we seen this key before?
	if existing, err := l.transfers.FindByIdempotencyKey(ctx, in.IdempotencyKey); err != nil {
		return nil, fmt.Errorf("idempotency lookup: %w", err)
	} else if existing != nil {
		return existing, nil // replay — return the original result
	}

	// 2. Build the domain object (this is where invariants are enforced).
	t, err := domain.NewTransfer(in.From, in.To, in.Money, in.IdempotencyKey)
	if err != nil {
		return nil, fmt.Errorf("invalid transfer: %w", err)
	}

	// 3. Do the work atomically.
	err = l.uow.WithTx(ctx, func(ctx context.Context) error {
		if err := t.Complete(); err != nil {
			return err
		}
		return l.transfers.Save(ctx, t)
	})

	if err != nil {
		return nil, err
	}

	return t, nil
}
