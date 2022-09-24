package logging

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Level is a log level.
type Level int32

const (
	// DebugLevel is a debug log level.
	DebugLevel Level = iota - 1
	// InfoLevel is an info log level, the default level.
	InfoLevel
	// WarnLevel is a warn log level.
	WarnLevel
	// ErrorLevel is a error log level.
	ErrorLevel
	// OffLevel is a log level that prevent logging.
	OffLevel
)

var (
	LevelName = map[Level]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		OffLevel:   "OFF",
	}
	LevelValue = map[string]Level{
		"DEBUG": DebugLevel,
		"INFO":  InfoLevel,
		"WARN":  WarnLevel,
		"ERROR": ErrorLevel,
		"OFF":   OffLevel,
	}
)

func (l Level) String() string {
	if name, ok := LevelName[l]; ok {
		return name
	}
	return "OFF"
}

// LevelEnabler decides whether a given logging level is enabled when logging a
// message.
type LevelEnabler interface {
	Enabled(Level) bool
}

// Enabled Each concrete Level value implements a static LevelEnabler which returns
// true for itself and all higher logging levels, except OffLevel.
// For example WarnLevel.Enabled() will return true for WarnLevel and ErrorLevel
// but return false for InfoLevel and DebugLevel, and OffLevel.
func (l Level) Enabled(proba Level) bool {
	if proba >= OffLevel {
		return false
	}
	return l <= proba
}

func (l Level) MarshalYAML() (any, error) {
	return l.String(), nil
}

func (l *Level) UnmarshalYAML(value *yaml.Node) error {
	v := value.Value
	v = strings.Trim(v, "\"'")
	v = strings.ToUpper(v)
	if byName, ok := LevelValue[v]; ok {
		*l = byName
		return nil
	}
	return fmt.Errorf("logging: unknown log level: %v", value.Value)
}
