package logging

// Logger is logger interface.
type Logger interface {
	// LevelEnabler enforce a Logger provide a Enabled method.
	LevelEnabler

	// Debug outputs a log at DebugLevel if the Logger enabled DebugLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the Print method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Debug(v ...any)
	// Debugln outputs a log at DebugLevel if the Logger enabled DebugLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Println method does but without an ending new line.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Debugln(v ...any)
	// Debugf outputs a log at DebugLevel if the Logger enabled DebugLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Printf method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Debugf(format string, v ...any)
	// Debugw outputs a log at DebugLevel if the Logger enabled DebugLevel.
	//
	// This method allow you log extra Fields to output.
	//
	// For fields with same name, usually, the later one override the former and builtin ones.
	// But a vendor may change this behavior by merge them.
	// Read the references of the vendor you chose about that.
	//
	// If you want format you msg field, do it before you call this method.
	// If formatting is expensive, you'd better call Enabled method before do that.
	Debugw(message string, field ...Field)

	// Info outputs a log at InfoLevel if the Logger enabled InfoLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the Print method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Info(v ...any)
	// Infoln outputs a log at InfoLevel if the Logger enabled InfoLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Println method does but without an ending new line.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Infoln(v ...any)
	// Infof outputs a log at InfoLevel if the Logger enabled InfoLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Printf method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Infof(format string, v ...any)
	// Infow outputs a log at InfoLevel if the Logger enabled InfoLevel.
	//
	// This method allow you log extra Fields to output.
	//
	// For fields with same name, usually, the later one override the former and builtin ones.
	// But a vendor may change this behavior by merge them.
	// Read the references of the vendor you chose about that.
	//
	// If you want format you msg field, do it before you call this method.
	// If formatting is expensive, you'd better call Enabled method before do that.
	Infow(message string, field ...Field)

	// Warn outputs a log at WarnLevel if the Logger enabled WarnLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the Print method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Warn(v ...any)
	// Warnln outputs a log at WarnLevel if the Logger enabled WarnLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Println method does but without an ending new line.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Warnln(v ...any)
	// Warnf outputs a log at WarnLevel if the Logger enabled WarnLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Printf method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Warnf(format string, v ...any)
	// Warnw outputs a log at WarnLevel if the Logger enabled WarnLevel.
	//
	// This method allow you log extra Fields to output.
	//
	// For fields with same name, usually, the later one override the former and builtin ones.
	// But a vendor may change this behavior by merge them.
	// Read the references of the vendor you chose about that.
	//
	// If you want format you msg field, do it before you call this method.
	// If formatting is expensive, you'd better call Enabled method before do that.
	Warnw(message string, field ...Field)

	// Error outputs a log at ErrorLevel if the Logger enabled ErrorLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the Print method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Error(v ...any)
	// Errorln outputs a log at ErrorLevel if the Logger enabled ErrorLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Println method does but without an ending new line.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Errorln(v ...any)
	// Errorf outputs a log at ErrorLevel if the Logger enabled ErrorLevel.
	//
	// The parameters would be used to forming the log msg field
	// via formatting like the fmt.Printf method does.
	//
	// If formatting is expensive, you'd better call Enabled method before do that.
	Errorf(format string, v ...any)
	// Errorw outputs a log at ErrorLevel if the Logger enabled ErrorLevel.
	//
	// This method allow you log extra Fields to output.
	//
	// For fields with same name, usually, the later one override the former and builtin ones.
	// But a vendor may change this behavior by merge them.
	// Read the references of the vendor you chose about that.
	//
	// If you want format you msg field, do it before you call this method.
	// If formatting is expensive, you'd better call Enabled method before do that.
	Errorw(message string, field ...Field)

	// WithField returns a new Logger which wrap extra Fields, including fields of current Logger.
	// This method should never return nil.
	//
	// For fields with same name, usually, the later one override the former and builtin ones.
	// But a vendor may change this behavior by merge them.
	// Read the references of the vendor you chose about that.
	WithField(field ...Field) Logger
}
