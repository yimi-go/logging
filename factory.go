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

type Factory interface {
	Logger(name string) Logger
}

func GetFactory() Factory {
	return factoryStore.Load().(*storedFactory).Factory
}

func SwapFactory(factory Factory) Factory {
	return factoryStore.Swap(&storedFactory{factory}).(*storedFactory).Factory
}
