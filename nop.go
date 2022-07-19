package logging

func NewNopLoggerFactory() Factory {
	return &nopLoggerFactory{}
}

type nopLoggerFactory struct{}

func (n *nopLoggerFactory) Logger(_ string) Logger { return &nopLogger{} }

type nopLogger struct{}

func (n *nopLogger) Enabled(_ Level) bool        { return false }
func (n *nopLogger) Debug(_ ...any)              {}
func (n *nopLogger) Debugln(_ ...any)            {}
func (n *nopLogger) Debugf(_ string, _ ...any)   {}
func (n *nopLogger) Debugw(_ string, _ ...Field) {}
func (n *nopLogger) Info(_ ...any)               {}
func (n *nopLogger) Infoln(_ ...any)             {}
func (n *nopLogger) Infof(_ string, _ ...any)    {}
func (n *nopLogger) Infow(_ string, _ ...Field)  {}
func (n *nopLogger) Warn(_ ...any)               {}
func (n *nopLogger) Warnln(_ ...any)             {}
func (n *nopLogger) Warnf(_ string, _ ...any)    {}
func (n *nopLogger) Warnw(_ string, _ ...Field)  {}
func (n *nopLogger) Error(_ ...any)              {}
func (n *nopLogger) Errorln(_ ...any)            {}
func (n *nopLogger) Errorf(_ string, _ ...any)   {}
func (n *nopLogger) Errorw(_ string, _ ...Field) {}
func (n *nopLogger) WithField(_ ...Field) Logger { return n }
