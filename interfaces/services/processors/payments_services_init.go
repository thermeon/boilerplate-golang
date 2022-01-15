package processors

import (
	"github.com/thermeon/boilerplate-golang/configs"
	"github.com/thermeon/boilerplate-golang/domain/abstractions"
)

// FactoryMap factory map type for processor factories.
type FactoryMap map[string]abstractions.Factory

// InitPaymentsAPIs Initialize the processors with factory map.
func InitPaymentsAPIs(c *configs.Config) (fm FactoryMap) {
	fm = make(map[string]abstractions.Factory)
	return
}
