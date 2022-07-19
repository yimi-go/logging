package hook

import (
	"github.com/yimi-go/logging"
)

type hooked struct {
	logging.Logger
	onLog func(meth string, param ...any)
}

func (l *hooked) Enabled(lvl logging.Level) bool {
	return l.Logger.Enabled(lvl)
}
func (l *hooked) Debug(v ...any) {
	if l.onLog != nil {
		l.onLog("Debug", v)
	}
	l.Logger.Debug(v...)
}
func (l *hooked) Debugln(v ...any) {
	if l.onLog != nil {
		l.onLog("Debugln", v)
	}
	l.Logger.Debugln(v...)
}
func (l *hooked) Debugf(format string, v ...any) {
	if l.onLog != nil {
		l.onLog("Debugf", format, v)
	}
	l.Logger.Debugf(format, v...)
}
func (l *hooked) Debugw(message string, field ...logging.Field) {
	if l.onLog != nil {
		l.onLog("Debugw", message, field)
	}
	l.Logger.Debugw(message, field...)
}
func (l *hooked) Info(v ...any) {
	if l.onLog != nil {
		l.onLog("Info", v)
	}
	l.Logger.Info(v...)
}
func (l *hooked) Infoln(v ...any) {
	if l.onLog != nil {
		l.onLog("Infoln", v)
	}
	l.Logger.Infoln(v...)
}
func (l *hooked) Infof(format string, v ...any) {
	if l.onLog != nil {
		l.onLog("Infof", format, v)
	}
	l.Logger.Infof(format, v...)
}
func (l *hooked) Infow(message string, field ...logging.Field) {
	if l.onLog != nil {
		l.onLog("Infow", message, field)
	}
	l.Logger.Infow(message, field...)
}
func (l *hooked) Warn(v ...any) {
	if l.onLog != nil {
		l.onLog("Warn", v)
	}
	l.Logger.Warn(v...)
}
func (l *hooked) Warnln(v ...any) {
	if l.onLog != nil {
		l.onLog("Warnln", v)
	}
	l.Logger.Warnln(v...)
}
func (l *hooked) Warnf(format string, v ...any) {
	if l.onLog != nil {
		l.onLog("Warnf", format, v)
	}
	l.Logger.Warnf(format, v...)
}
func (l *hooked) Warnw(message string, field ...logging.Field) {
	if l.onLog != nil {
		l.onLog("Warnw", message, field)
	}
	l.Logger.Warnw(message, field...)
}
func (l *hooked) Error(v ...any) {
	if l.onLog != nil {
		l.onLog("Error", v)
	}
	l.Logger.Error(v...)
}
func (l *hooked) Errorln(v ...any) {
	if l.onLog != nil {
		l.onLog("Errorln", v)
	}
	l.Logger.Errorln(v...)
}
func (l *hooked) Errorf(format string, v ...any) {
	if l.onLog != nil {
		l.onLog("Errorf", format, v)
	}
	l.Logger.Errorf(format, v...)
}
func (l *hooked) Errorw(message string, field ...logging.Field) {
	if l.onLog != nil {
		l.onLog("Errorw", message, field)
	}
	l.Logger.Errorw(message, field...)
}
func (l *hooked) WithField(field ...logging.Field) logging.Logger {
	return &hooked{
		Logger: l.Logger.WithField(field...),
		onLog:  l.onLog,
	}
}

type hookedFactory struct {
	factory logging.Factory
	onLog   func(meth string, param ...any)
}

func (f *hookedFactory) Logger(name string) logging.Logger {
	return &hooked{
		Logger: f.factory.Logger(name),
		onLog:  f.onLog,
	}
}

func Hooked(factory logging.Factory, onLog func(meth string, param ...any)) logging.Factory {
	return &hookedFactory{
		factory: factory,
		onLog:   onLog,
	}
}
