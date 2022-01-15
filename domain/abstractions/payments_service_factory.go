package abstractions

type ProcessorName string
type API string

type Selector interface {
	Select(p ProcessorName, api API) (factory Factory, err error)
}

type Factory interface {
	GetLaneService() LaneService
	GetTransactionService() TransactionService
}
