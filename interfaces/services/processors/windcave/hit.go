package windcave

import (
	"context"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type hitAPI struct {
	conf *configs.Config
}

func NewHitAPI(c *configs.Config) abstractions.TransactionService {
	return &hitAPI{
		conf: c,
	}
}

func (e hitAPI) Authorization(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Sale(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Capture(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Option(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Refund(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Status(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Void(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Tokenize(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e hitAPI) Name() (name string) {
	panic("implement me")
}

func (e hitAPI) LogName() (logName string) {
	panic("implement me")
}
