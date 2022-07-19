package logging

import (
	"reflect"
	"testing"
)

func TestGetFactory(t *testing.T) {
	factory := GetFactory()
	if factory == nil {
		t.Errorf("want a factory, got nil")
	}
	if !reflect.DeepEqual(factory, factoryStore.Load().(*storedFactory).Factory) {
		t.Errorf("got %v, want %v", factory, factoryStore.Load().(*storedFactory).Factory)
	}
}

type tFactory struct {
	nopLoggerFactory
}

func TestSwapFactory(t *testing.T) {
	origin := GetFactory()
	defer func() {
		factoryStore.Store(&storedFactory{origin})
	}()
	newFactory := &tFactory{}
	swap := SwapFactory(newFactory)
	if !reflect.DeepEqual(origin, swap) {
		t.Errorf("swap: %v, origin: %v", swap, origin)
	}
	factory := GetFactory()
	if reflect.DeepEqual(factory, origin) {
		t.Errorf("swap not work")
	}
	if !reflect.DeepEqual(newFactory, factory) {
		t.Errorf("swap not work")
	}
}
