package logging

import (
	"go.uber.org/atomic"
)

func init() {
	factoryStore.Store(&storedFactory{&nopLoggerFactory{}})
}

var factoryStore atomic.Value

type storedFactory struct {
	Factory
}

// Factory is a Factory that produces Logger by name.
type Factory interface {
	// Logger returns a Logger by name, which may be cached.
	// It should never return nil.
	Logger(name string) Logger
}

// GetFactory returns the registered Factory. It should never return nil.
// By default, a nop Logger Factory will be returned.
//
// For any production projects, a vendor provided Factory should be registered
// first via SwapFactory before first calling GetFactory
func GetFactory() Factory {
	return factoryStore.Load().(*storedFactory).Factory
}

// SwapFactory registers new Factory, and returns the origin Factory.
//
// For any production projects, a vendor provided Factory should be registered
// first via SwapFactory before first calling GetFactory
func SwapFactory(factory Factory) Factory {
	return factoryStore.Swap(&storedFactory{factory}).(*storedFactory).Factory
}
