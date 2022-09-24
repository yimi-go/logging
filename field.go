package logging

import (
	"context"
	"fmt"
	"time"
)

type FieldType uint8

const (
	// UnknownType is the default field type.
	UnknownType FieldType = iota
	// BinaryType indicates that the field carries an opaque binary blob.
	BinaryType
	// BoolType indicates that the field carries a bool.
	BoolType
	// Complex128Type indicates that the field carries a complex128.
	Complex128Type
	// Complex64Type indicates that the field carries a complex128.
	Complex64Type
	// DurationType indicates that the field carries a time.Duration.
	DurationType
	// Float64Type indicates that the field carries a float64.
	Float64Type
	// Float32Type indicates that the field carries a float32.
	Float32Type
	// Int64Type indicates that the field carries an int64.
	Int64Type
	// Int32Type indicates that the field carries an int32.
	Int32Type
	// Int16Type indicates that the field carries an int16.
	Int16Type
	// Int8Type indicates that the field carries an int8.
	Int8Type
	// StringType indicates that the field carries a string.
	StringType
	// TimeType indicates that the field carries a time.Time.
	TimeType
	// Uint64Type indicates that the field carries an uint64.
	Uint64Type
	// Uint32Type indicates that the field carries an uint32.
	Uint32Type
	// Uint16Type indicates that the field carries an uint16.
	Uint16Type
	// Uint8Type indicates that the field carries an uint8.
	Uint8Type
	// UintptrType indicates that the field carries an uintptr.
	UintptrType
	// StringerType indicates that the field carries a fmt.Stringer.
	StringerType
	// ErrorType indicates that the field carries an error.
	ErrorType
	// StackType indicates that the field captures stacktrace of the current goroutine.
	StackType
)

// Field is logging field.
type Field interface {
	// Key is field key.
	Key() string
	// Type is field type.
	Type() FieldType
	// Value is field value.
	Value() any
}

type field struct {
	key string
	typ FieldType
	val any
}

func (f field) Key() string     { return f.key }
func (f field) Type() FieldType { return f.typ }
func (f field) Value() any      { return f.val }

// Any creates a Field of UnknownType value.
func Any(key string, value any) Field {
	return field{key: key, typ: UnknownType, val: value}
}

// Binary creates a Field of BinaryType value.
func Binary(key string, value []byte) Field {
	val := make([]byte, len(value))
	copy(val, value)
	return field{key: key, typ: BinaryType, val: val}
}

// Bool creates a Field of BoolType value.
func Bool(key string, value bool) Field {
	return field{key: key, typ: BoolType, val: value}
}

// Boolp creates a Field of BoolType value if value is not nil, or a Field of UnknownType value if value is nil.
func Boolp(key string, value *bool) Field {
	if value == nil {
		return Any(key, value)
	}
	return Bool(key, *value)
}

// Complex128 creates a Field of Complex128Type value.
func Complex128(key string, value complex128) Field {
	return field{key: key, typ: Complex128Type, val: value}
}

// Complex128p creates a Field of Complex128Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Complex128p(key string, value *complex128) Field {
	if value == nil {
		return Any(key, value)
	}
	return Complex128(key, *value)
}

// Complex64 creates a Field of Complex64Type value.
func Complex64(key string, value complex64) Field {
	return field{key: key, typ: Complex64Type, val: value}
}

// Complex64p creates a Field of Complex64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Complex64p(key string, value *complex64) Field {
	if value == nil {
		return Any(key, value)
	}
	return Complex64(key, *value)
}

// Duration creates a Field of DurationType value.
func Duration(key string, value time.Duration) Field {
	return field{key: key, typ: DurationType, val: value}
}

// Durationp creates a Field of DurationType value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Durationp(key string, value *time.Duration) Field {
	if value == nil {
		return Any(key, value)
	}
	return Duration(key, *value)
}

// Float64 creates a Field of Float64Type value.
func Float64(key string, value float64) Field {
	return field{key: key, typ: Float64Type, val: value}
}

// Float64p creates a Field of Float64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Float64p(key string, value *float64) Field {
	if value == nil {
		return Any(key, value)
	}
	return Float64(key, *value)
}

// Float32 creates a Field of Float32Type value.
func Float32(key string, value float32) Field {
	return field{key: key, typ: Float32Type, val: value}
}

// Float32p creates a Field of Float32Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Float32p(key string, value *float32) Field {
	if value == nil {
		return Any(key, value)
	}
	return Float32(key, *value)
}

// Int creates a Field of Int64Type value.
func Int(key string, value int) Field {
	return field{key: key, typ: Int64Type, val: int64(value)}
}

// Intp creates a Field of Int64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Intp(key string, value *int) Field {
	if value == nil {
		return Any(key, (*int64)(nil))
	}
	return Int(key, *value)
}

// Int64 creates a Field of Int64Type value.
func Int64(key string, value int64) Field {
	return field{key: key, typ: Int64Type, val: value}
}

// Int64p creates a Field of Int64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Int64p(key string, value *int64) Field {
	if value == nil {
		return Any(key, value)
	}
	return Int64(key, *value)
}

// Int32 creates a Field of Int32Type value.
func Int32(key string, value int32) Field {
	return field{key: key, typ: Int32Type, val: value}
}

// Int32p creates a Field of Int32Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Int32p(key string, value *int32) Field {
	if value == nil {
		return Any(key, value)
	}
	return Int32(key, *value)
}

// Int16 creates a Field of Int16Type value.
func Int16(key string, value int16) Field {
	return field{key: key, typ: Int16Type, val: value}
}

// Int16p creates a Field of Int16Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Int16p(key string, value *int16) Field {
	if value == nil {
		return Any(key, value)
	}
	return Int16(key, *value)
}

// Int8 creates a Field of Int8Type value.
func Int8(key string, value int8) Field {
	return field{key: key, typ: Int8Type, val: value}
}

// Int8p creates a Field of Int8Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Int8p(key string, value *int8) Field {
	if value == nil {
		return Any(key, value)
	}
	return Int8(key, *value)
}

// String creates a Field of StringType value.
func String(key string, value string) Field {
	return field{key: key, typ: StringType, val: value}
}

// Stringp creates a Field of StringType value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Stringp(key string, value *string) Field {
	if value == nil {
		return Any(key, value)
	}
	return String(key, *value)
}

// Time creates a Field of TimeType value.
func Time(key string, value time.Time) Field {
	return field{key: key, typ: TimeType, val: value}
}

// Timep creates a Field of TimeType value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Timep(key string, value *time.Time) Field {
	if value == nil {
		return Any(key, value)
	}
	return Time(key, *value)
}

// Uint creates a Field of Uint64Type value.
func Uint(key string, value uint) Field {
	return field{key: key, typ: Uint64Type, val: uint64(value)}
}

// Uintp creates a Field of Uint64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uintp(key string, value *uint) Field {
	if value == nil {
		return Any(key, (*uint64)(nil))
	}
	return Uint(key, *value)
}

// Uint64 creates a Field of Uint64Type value.
func Uint64(key string, value uint64) Field {
	return field{key: key, typ: Uint64Type, val: value}
}

// Uint64p creates a Field of Uint64Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uint64p(key string, value *uint64) Field {
	if value == nil {
		return Any(key, value)
	}
	return Uint64(key, *value)
}

// Uint32 creates a Field of Uint32Type value.
func Uint32(key string, value uint32) Field {
	return field{key: key, typ: Uint32Type, val: value}
}

// Uint32p creates a Field of Uint32Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uint32p(key string, value *uint32) Field {
	if value == nil {
		return Any(key, value)
	}
	return Uint32(key, *value)
}

// Uint16 creates a Field of Uint16Type value.
func Uint16(key string, value uint16) Field {
	return field{key: key, typ: Uint16Type, val: value}
}

// Uint16p creates a Field of Uint16Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uint16p(key string, value *uint16) Field {
	if value == nil {
		return Any(key, value)
	}
	return Uint16(key, *value)
}

// Uint8 creates a Field of Uint8Type value.
func Uint8(key string, value uint8) Field {
	return field{key: key, typ: Uint8Type, val: value}
}

// Uint8p creates a Field of Uint8Type value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uint8p(key string, value *uint8) Field {
	if value == nil {
		return Any(key, value)
	}
	return Uint8(key, *value)
}

// Uintptr creates a Field of UintptrType value.
func Uintptr(key string, value uintptr) Field {
	return field{key: key, typ: UintptrType, val: value}
}

// Uintptrp creates a Field of UintptrType value if value is not nil,
// or a Field of UnknownType value if value is nil.
func Uintptrp(key string, value *uintptr) Field {
	if value == nil {
		return Any(key, value)
	}
	return Uintptr(key, *value)
}

// Stringer creates a Field of StringerType value.
func Stringer(key string, value fmt.Stringer) Field {
	return field{key: key, typ: StringerType, val: value}
}

// Error creates a Field of ErrorType value, with key "error".
func Error(err error) Field {
	return field{key: "error", typ: ErrorType, val: err}
}

// NamedError create a Field of ErrorType value.
func NamedError(key string, err error) Field {
	return field{key: key, typ: ErrorType, val: err}
}

// Stack create a Field that captures stacktrace of the current goroutine.
func Stack(key string) Field {
	return field{key: key, typ: StackType, val: 0}
}

// StackSkip create a Field that captures stacktrace of the current goroutine,
// skipping the given number of frames from the top of the stacktrace.
func StackSkip(key string, skip int) Field {
	return field{key: key, typ: StackType, val: skip}
}

type fieldKey struct{}

// NewContext wraps fields into a new context and return it.
func NewContext(ctx context.Context, field ...Field) context.Context {
	return context.WithValue(ctx, fieldKey{}, field)
}

// WithContextField tries extract fields from context and returns a Logger with these fields.
// If no fields found, the origin logger is returned.
func WithContextField(ctx context.Context, logger Logger) Logger {
	fields, ok := ctx.Value(fieldKey{}).([]Field)
	if !ok {
		return logger
	}
	if len(fields) == 0 {
		return logger
	}
	return logger.WithField(fields...)
}
