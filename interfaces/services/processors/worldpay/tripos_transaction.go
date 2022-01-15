package worldpay

import (
	"context"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
	"github.com/thermeon/boilerplate-golang/domain/entities"
)

type triPosAPI struct {
}

func NewTriPosAPI() abstractions.TransactionService {
	return &triPosAPI{}
}

func (t triPosAPI) Authorization(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Sale(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Capture(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Option(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Refund(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Status(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Void(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Tokenize(ctx context.Context, txn *entities.Transaction) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) Name() (name string) {
	//TODO implement me
	panic("implement me")
}

func (t triPosAPI) LogName() (logName string) {
	//TODO implement me
	panic("implement me")
}
