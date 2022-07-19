package logging

type Logger interface {
	LevelEnabler

	Debug(v ...any)
	Debugln(v ...any)
	Debugf(format string, v ...any)
	Debugw(message string, field ...Field)

	Info(v ...any)
	Infoln(v ...any)
	Infof(format string, v ...any)
	Infow(message string, field ...Field)

	Warn(v ...any)
	Warnln(v ...any)
	Warnf(format string, v ...any)
	Warnw(message string, field ...Field)

	Error(v ...any)
	Errorln(v ...any)
	Errorf(format string, v ...any)
	Errorw(message string, field ...Field)

	WithField(field ...Field) Logger
}
