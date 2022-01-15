package abstractions

import (
	"context"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type TransactionService interface {
	Authorization(ctx context.Context, txn *entities.Transaction) (err error)
	Sale(ctx context.Context, txn *entities.Transaction) (err error)
	Capture(ctx context.Context, txn *entities.Transaction) (err error)
	Option(ctx context.Context, txn *entities.Transaction) (err error)
	Refund(ctx context.Context, txn *entities.Transaction) (err error)
	Status(ctx context.Context, txn *entities.Transaction) (err error)
	Void(ctx context.Context, txn *entities.Transaction) (err error)
	Tokenize(ctx context.Context, txn *entities.Transaction) (err error)
	Name() (name string)
	LogName() (logName string)
}
