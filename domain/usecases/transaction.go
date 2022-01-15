package usecases

import (
	"context"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
	"github.com/thermeon/boilerplate-golang/interfaces/services/processors"
	"github.com/thermeon/boilerplate-golang/log"
)

type transactionUseCases struct {
	abstractions.Container
}

func NewTransactionUseCases(c abstractions.Container) transactionUseCases {
	return transactionUseCases{
		Container: c,
	}
}

func (t *transactionUseCases) Authorization(ctx context.Context, txn entities.Transaction) (err error) {
	_, err = t.BackendSelector.Select(processors.Windcave, processors.Express)
	if err != nil {
		log.Error(ctx, "Unable to select processor. Error: %s", err)
		return
	}
	return
}
