package windcave

import (
	"context"
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type postAPI struct {
	conf *configs.Config
}

func NewPostAPI(c *configs.Config) abstractions.TransactionService {
	return &postAPI{conf: c}
}

func (e postAPI) Authorization(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Sale(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Capture(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Option(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Refund(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Status(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Void(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Tokenize(ctx context.Context, txn *entities.Transaction) (err error) {
	panic("implement me")
}

func (e postAPI) Name() (name string) {
	panic("implement me")
}

func (e postAPI) LogName() (logName string) {
	panic("implement me")
}
