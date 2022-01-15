package worldpay

import (
	"context"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type expressAPI struct {
	conf *configs.Config
}

func NewExpressAPI(c *configs.Config) abstractions.TransactionService {
	return &expressAPI{
		conf: c,
	}
}

func (e expressAPI) Authorization(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Sale(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Capture(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Option(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Refund(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Status(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Void(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Tokenize(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e expressAPI) Name() (name string) {
	panic("implement me")
}

func (e expressAPI) LogName() (logName string) {
	panic("implement me")
}
