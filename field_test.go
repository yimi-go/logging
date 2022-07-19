package logging

import (
	"context"
	"io"
	"reflect"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type tl struct {
	nopLogger
	field []Field
}

func (t *tl) WithField(field ...Field) Logger {
	return &tl{field: append(t.field, field...)}
}

type hanBoolStringer bool

func (h hanBoolStringer) String() string {
	switch bool(h) {
	case true:
		return "真"
	case false:
		return "假"
	}
	panic("no other value")
}

func TestField_Key(t *testing.T) {
	f := field{key: "abc"}
	assert.Equal(t, "abc", f.Key())
}

func TestField_Type(t *testing.T) {
	f := field{typ: StringType}
	assert.Equal(t, StringType, f.Type())
}

func TestField_Value(t *testing.T) {
	f := field{val: 123}
	assert.Equal(t, 123, f.Value())
}

func TestAny(t *testing.T) {
	key, value := "key", map[string]any{"a": 1}
	f := Any(key, value)
	assert.Equal(t, field{key, UnknownType, value}, f)
}

func TestBinary(t *testing.T) {
	key, value := "key", []byte{'a', 'b', 'c'}
	f := Binary(key, value)
	assert.Equal(t, field{key, BinaryType, value}, f)
}

func TestBool(t *testing.T) {
	key := "key"
	f := Bool(key, true)
	assert.Equal(t, field{key, BoolType, true}, f)
}

func TestBoolp(t *testing.T) {
	v := true
	type args struct {
		key   string
		value *bool
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*bool)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", BoolType, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Boolp(tt.args.key, tt.args.value), "Boolp(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestComplex128(t *testing.T) {
	key := "key"
	f := Complex128(key, 1+2i)
	assert.Equal(t, field{key, Complex128Type, 1 + 2i}, f)
}

func TestComplex128p(t *testing.T) {
	v := 1 + 2i
	type args struct {
		key   string
		value *complex128
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*complex128)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Complex128Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Complex128p(tt.args.key, tt.args.value),
				"Complex128p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestComplex64(t *testing.T) {
	key := "key"
	f := Complex64(key, (complex64)(1+2i))
	assert.Equal(t, field{key, Complex64Type, (complex64)(1 + 2i)}, f)
}

func TestComplex64p(t *testing.T) {
	v := (complex64)(1 + 2i)
	type args struct {
		key   string
		value *complex64
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*complex64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Complex64Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Complex64p(tt.args.key, tt.args.value),
				"Complex64p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestDuration(t *testing.T) {
	key := "key"
	f := Duration(key, time.Second)
	assert.Equal(t, field{key, DurationType, time.Second}, f)
}

func TestDurationp(t *testing.T) {
	v := time.Second
	type args struct {
		key   string
		value *time.Duration
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*time.Duration)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", DurationType, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Durationp(tt.args.key, tt.args.value),
				"Durationp(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestFloat64(t *testing.T) {
	key := "key"
	f := Float64(key, 1.2)
	assert.Equal(t, field{key, Float64Type, 1.2}, f)
}

func TestFloat64p(t *testing.T) {
	v := 1.2
	type args struct {
		key   string
		value *float64
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*float64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Float64Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Float64p(tt.args.key, tt.args.value),
				"Float64p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestFloat32(t *testing.T) {
	key := "key"
	f := Float32(key, float32(1.2))
	assert.Equal(t, field{key, Float32Type, float32(1.2)}, f)
}

func TestFloat32p(t *testing.T) {
	v := float32(1.2)
	type args struct {
		key   string
		value *float32
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*float32)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Float32Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Float32p(tt.args.key, tt.args.value),
				"Float32p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestInt(t *testing.T) {
	key := "key"
	f := Int(key, 2)
	assert.Equal(t, field{key, Int64Type, int64(2)}, f)
}

func TestIntp(t *testing.T) {
	v := 2
	type args struct {
		key   string
		value *int
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*int64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Int64Type, int64(v)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Intp(tt.args.key, tt.args.value), "Intp(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestInt64(t *testing.T) {
	key := "key"
	f := Int64(key, int64(2))
	assert.Equal(t, field{key, Int64Type, int64(2)}, f)
}

func TestInt64p(t *testing.T) {
	v := int64(2)
	type args struct {
		key   string
		value *int64
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*int64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Int64Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int64p(tt.args.key, tt.args.value), "Int64p(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestInt32(t *testing.T) {
	key := "key"
	f := Int32(key, int32(2))
	assert.Equal(t, field{key, Int32Type, int32(2)}, f)
}

func TestInt32p(t *testing.T) {
	v := int32(2)
	type args struct {
		key   string
		value *int32
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*int32)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Int32Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int32p(tt.args.key, tt.args.value), "Int32p(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestInt16(t *testing.T) {
	key := "key"
	f := Int16(key, int16(2))
	assert.Equal(t, field{key, Int16Type, int16(2)}, f)
}

func TestInt16p(t *testing.T) {
	v := int16(2)
	type args struct {
		key   string
		value *int16
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*int16)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Int16Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int16p(tt.args.key, tt.args.value), "Int16p(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestInt8(t *testing.T) {
	key := "key"
	f := Int8(key, int8(2))
	assert.Equal(t, field{key, Int8Type, int8(2)}, f)
}

func TestInt8p(t *testing.T) {
	v := int8(2)
	type args struct {
		key   string
		value *int8
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*int8)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Int8Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Int8p(tt.args.key, tt.args.value), "Int8p(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestString(t *testing.T) {
	key := "key"
	f := String(key, "val")
	assert.Equal(t, field{key, StringType, "val"}, f)
}

func TestStringp(t *testing.T) {
	v := "val"
	type args struct {
		key   string
		value *string
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*string)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", StringType, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Stringp(tt.args.key, tt.args.value),
				"Stringp(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestTime(t *testing.T) {
	key := "key"
	val := time.Now()
	f := Time(key, val)
	assert.Equal(t, field{key, TimeType, val}, f)
}

func TestTimep(t *testing.T) {
	v := time.Now()
	type args struct {
		key   string
		value *time.Time
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*time.Time)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", TimeType, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Timep(tt.args.key, tt.args.value), "Timep(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestWithFieldContext(t *testing.T) {
	field := sf("foo", "bar")
	ctx := NewContext(context.Background(), field)
	fields, ok := ctx.Value(fieldKey{}).([]Field)
	if !ok {
		t.Fatalf("expect ok, but not")
	}
	if len(fields) != 1 {
		t.Fatalf("expect len 1, got %v", len(fields))
	}
	if !reflect.DeepEqual(field, fields[0]) {
		t.Errorf("want %v, got %v", field, fields[0])
	}
}

func TestUint(t *testing.T) {
	key := "key"
	f := Uint(key, uint(2))
	assert.Equal(t, field{key, Uint64Type, uint64(2)}, f)
}

func TestUintp(t *testing.T) {
	v := uint(2)
	type args struct {
		key   string
		value *uint
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uint64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Uint64Type, uint64(v)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Uintp(tt.args.key, tt.args.value), "Uintp(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestUint64(t *testing.T) {
	key := "key"
	f := Uint64(key, uint64(2))
	assert.Equal(t, field{key, Uint64Type, uint64(2)}, f)
}

func TestUint64p(t *testing.T) {
	v := uint64(2)
	type args struct {
		key   string
		value *uint64
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uint64)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Uint64Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Uint64p(tt.args.key, tt.args.value),
				"Uint64p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestUint32(t *testing.T) {
	key := "key"
	f := Uint32(key, uint32(2))
	assert.Equal(t, field{key, Uint32Type, uint32(2)}, f)
}

func TestUint32p(t *testing.T) {
	v := uint32(2)
	type args struct {
		key   string
		value *uint32
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uint32)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Uint32Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Uint32p(tt.args.key, tt.args.value),
				"Uint32p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestUint16(t *testing.T) {
	key := "key"
	f := Uint16(key, uint16(2))
	assert.Equal(t, field{key, Uint16Type, uint16(2)}, f)
}

func TestUint16p(t *testing.T) {
	v := uint16(2)
	type args struct {
		key   string
		value *uint16
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uint16)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Uint16Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Uint16p(tt.args.key, tt.args.value),
				"Uint16p(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestUint8(t *testing.T) {
	key := "key"
	f := Uint8(key, uint8(2))
	assert.Equal(t, field{key, Uint8Type, uint8(2)}, f)
}

func TestUint8p(t *testing.T) {
	v := uint8(2)
	type args struct {
		key   string
		value *uint8
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uint8)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", Uint8Type, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Uint8p(tt.args.key, tt.args.value), "Uint8p(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestUintptr(t *testing.T) {
	key := "key"
	val := (uintptr)(unsafe.Pointer(&key))
	f := Uintptr(key, val)
	assert.Equal(t, field{key, UintptrType, val}, f)
}

func TestUintptrp(t *testing.T) {
	s := "some"
	v := (uintptr)(unsafe.Pointer(&s))
	type args struct {
		key   string
		value *uintptr
	}
	tests := []struct {
		name string
		args args
		want Field
	}{
		{
			name: "nil",
			args: args{"key", nil},
			want: field{"key", UnknownType, (*uintptr)(nil)},
		},
		{
			name: "v",
			args: args{"key", &v},
			want: field{"key", UintptrType, v},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				Uintptrp(tt.args.key, tt.args.value),
				"Uintptrp(%v, %v)",
				tt.args.key,
				tt.args.value,
			)
		})
	}
}

func TestStringer(t *testing.T) {
	key := "key"
	val := hanBoolStringer(true)
	f := Stringer(key, val)
	assert.Equal(t, field{key, StringerType, val}, f)
}

func TestError(t *testing.T) {
	val := io.EOF
	f := Error(val)
	assert.Equal(t, field{"error", ErrorType, val}, f)
}

func TestNamedError(t *testing.T) {
	key := "key"
	val := io.EOF
	f := NamedError(key, val)
	assert.Equal(t, field{key, ErrorType, val}, f)
}

func TestStack(t *testing.T) {
	key := "key"
	f := Stack(key)
	assert.Equal(t, field{"key", StackType, 0}, f)
}

func TestStackSkip(t *testing.T) {
	key := "key"
	f := StackSkip(key, 3)
	assert.Equal(t, field{"key", StackType, 3}, f)
}

func TestWithContextField(t *testing.T) {
	t.Run("not_set", func(t *testing.T) {
		l := &tl{}
		nl := WithContextField(context.Background(), l).(*tl)
		if len(nl.field) != 0 {
			t.Errorf("want len 0, got %v", len(nl.field))
		}
	})
	t.Run("empty_field", func(t *testing.T) {
		l := &tl{}
		ctx := NewContext(context.Background())
		nl := WithContextField(ctx, l).(*tl)
		if len(nl.field) != 0 {
			t.Errorf("want len 0, got %v", len(nl.field))
		}
	})
	t.Run("fields", func(t *testing.T) {
		l := &tl{}
		ctx := NewContext(context.Background(), sf("foo", "bar"))
		nl := WithContextField(ctx, l).(*tl)
		if len(nl.field) == 0 {
			t.Errorf("want len 1, got %v", len(nl.field))
		}
	})
}
