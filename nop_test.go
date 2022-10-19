package logging

import (
	"reflect"
	"testing"
)

func sf(key, value string) Field {
	return field{value, key, StringType}
}

func TestNopLogger(t *testing.T) {
	l := nopLogger{}
	for i := DebugLevel; i <= OffLevel; i++ {
		if l.Enabled(i) {
			t.Errorf("unexpected enabled: %d", i)
			return
		}
	}
	l.Debug("test")
	l.Debugln("test")
	l.Debugf("%v", "test")
	l.Debugw("test", sf("key", "value"))
	l.Info("test")
	l.Infoln("test")
	l.Infof("%v", "test")
	l.Infow("test", sf("key", "value"))
	l.Warn("test")
	l.Warnln("test")
	l.Warnf("%v", "test")
	l.Warnw("test", sf("key", "value"))
	l.Error("test")
	l.Errorln("test")
	l.Errorf("%v", "test")
	l.Errorw("test", sf("key", "value"))
	nl := l.WithField(sf("key", "value"))
	if nl == nil {
		t.Errorf("unexpected nil")
		return
	}
}

func Test_nopFactory_Logger(t *testing.T) {
	f := &nopLoggerFactory{}
	if !reflect.DeepEqual(f.Logger("any"), &nopLogger{}) {
		t.Errorf("Logger() = %v, want %v", f.Logger("any"), &nopLogger{})
	}
}

func TestNewNopLoggerFactory(t *testing.T) {
	got := NewNopLoggerFactory()
	want := &nopLoggerFactory{}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
