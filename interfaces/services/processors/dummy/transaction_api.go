package dummy

import (
	"context"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type dummyAPI struct {
	conf *configs.Config
}

func NewDummyAPI(c *configs.Config) abstractions.TransactionService {
	return &dummyAPI{conf: c}
}

func (e dummyAPI) Authorization(ctx context.Context, txn *entities.Transaction) (err error) {
	if txn.IsCardholderNotPresent() {
		return e.CardholderNotPresentAuthorization(ctx, txn)
	}
	return
}

func (e dummyAPI) Sale(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Capture(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Option(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Refund(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Status(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Void(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Tokenize(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e dummyAPI) Name() (name string) {
	panic("implement me")
}

func (e dummyAPI) LogName() (logName string) {
	panic("implement me")
}

func (e dummyAPI) CardholderNotPresentAuthorization(ctx context.Context, txn *entities.Transaction) (err error) {
	return
}
