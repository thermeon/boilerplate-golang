package processors

import (
	"fmt"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
)

const (
	Dummy abstractions.ProcessorName = "dummy"
)

type selector struct {
	fm FactoryMap
}

// NewPaymentServiceSelector initialized the processor selector.
func NewPaymentServiceSelector(factoryMap FactoryMap) abstractions.Selector {
	return &selector{
		fm: factoryMap,
	}
}

// Select get processor factory instances by processor name and API name.
func (s *selector) Select(p abstractions.ProcessorName, api abstractions.API) (fac abstractions.Factory, err error) {
	return
}

// getFactoryMapKey generate the factory map key based on processor name and API name.
func getFactoryMapKey(p abstractions.ProcessorName, api abstractions.API) string {
	return fmt.Sprintf("%s_%s", p, api)
}
