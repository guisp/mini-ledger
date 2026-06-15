package domain

import (
	"errors"
	"fmt"
	"time"
)

type Transfer struct {
	ID             string         `json:"id"`
	From           AccountID      `json:"from"`
	To             AccountID      `json:"to"`
	Money          Money          `json:"money"`
	Status         TransferStatus `json:"status"`
	IdempotencyKey string         `json:"idempotency_key"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

func NewTransfer(from, to AccountID, money Money, idempotencyKey string) (*Transfer, error) {
	if from == "" || to == "" {
		return nil, errors.New("from and to accounts are required")
	}
	if from == to {
		return nil, errors.New("cannot transfer to the same account")
	}
	if money.Amount <= 0 {
		return nil, errors.New("transfer amount must be positive")
	}
	if money.Currency == "" {
		return nil, errors.New("currency is required")
	}
	if idempotencyKey == "" {
		return nil, errors.New("idempotency key is required")
	}

	now := time.Now().UTC()
	return &Transfer{
		From:           from,
		To:             to,
		Money:          money,
		Status:         TransferPending,
		IdempotencyKey: idempotencyKey,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (t *Transfer) Complete() error {
	if t.Status != TransferPending {
		return fmt.Errorf("cannot complete transfer in state %s", t.Status)
	}
	t.Status = TransferCompleted
	t.UpdatedAt = time.Now().UTC()
	return nil
}
