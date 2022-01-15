package processors

import (
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
)

type factory struct {
	abstractions.LaneService
	abstractions.TransactionService
}

func NewFactory(lane abstractions.LaneService, tns abstractions.TransactionService) abstractions.Factory {
	return &factory{
		LaneService:        lane,
		TransactionService: tns,
	}
}

func (f factory) GetLaneService() abstractions.LaneService {
	return f.LaneService
}

func (f factory) GetTransactionService() abstractions.TransactionService {
	return f.TransactionService
}
