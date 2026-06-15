package ledger

type Ledger struct {
	transfers TransferStore
	accounts  AccountStore
	uow       UnitOfWork
}

func NewLedger(t TransferStore, a AccountStore, uow UnitOfWork) *Ledger {
	return &Ledger{transfers: t, accounts: a, uow: uow}
}
