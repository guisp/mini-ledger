package main

import (
	"fmt"
	"time"

	"github.com/guisp/miniledger/internal/domain"
)

func main() {
	transfer := domain.Transfer{
		ID:             "12345",
		From:           "account1",
		To:             "account2",
		Money:          domain.Money{Amount: 1000, Currency: "USD"},
		Status:         domain.TransferPending,
		IdempotencyKey: "abcde",
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	fmt.Println("Hello, World!", transfer.ID)
}
