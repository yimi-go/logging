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
	assert.Equal(t, field{value, key, UnknownType}, f)
}

func TestBinary(t *testing.T) {
	key, value := "key", []byte{'a', 'b', 'c'}
	f := Binary(key, value)
	assert.Equal(t, field{value, key, BinaryType}, f)
}

func TestBool(t *testing.T) {
	key := "key"
	f := Bool(key, true)
	assert.Equal(t, field{true, key, BoolType}, f)
}

func TestBoolp(t *testing.T) {
	v := true
	type args struct {
		value *bool
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*bool)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", BoolType},
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
	assert.Equal(t, field{1 + 2i, key, Complex128Type}, f)
}

func TestComplex128p(t *testing.T) {
	v := 1 + 2i
	type args struct {
		value *complex128
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*complex128)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Complex128Type},
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
	assert.Equal(t, field{(complex64)(1 + 2i), key, Complex64Type}, f)
}

func TestComplex64p(t *testing.T) {
	v := (complex64)(1 + 2i)
	type args struct {
		value *complex64
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*complex64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Complex64Type},
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
	assert.Equal(t, field{time.Second, key, DurationType}, f)
}

func TestDurationp(t *testing.T) {
	v := time.Second
	type args struct {
		value *time.Duration
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*time.Duration)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", DurationType},
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
	assert.Equal(t, field{1.2, key, Float64Type}, f)
}

func TestFloat64p(t *testing.T) {
	v := 1.2
	type args struct {
		value *float64
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*float64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Float64Type},
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
	assert.Equal(t, field{float32(1.2), key, Float32Type}, f)
}

func TestFloat32p(t *testing.T) {
	v := float32(1.2)
	type args struct {
		value *float32
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*float32)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Float32Type},
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
	assert.Equal(t, field{int64(2), key, Int64Type}, f)
}

func TestIntp(t *testing.T) {
	v := 2
	type args struct {
		value *int
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*int64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{int64(v), "key", Int64Type},
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
	assert.Equal(t, field{int64(2), key, Int64Type}, f)
}

func TestInt64p(t *testing.T) {
	v := int64(2)
	type args struct {
		value *int64
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*int64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Int64Type},
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
	assert.Equal(t, field{int32(2), key, Int32Type}, f)
}

func TestInt32p(t *testing.T) {
	v := int32(2)
	type args struct {
		value *int32
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*int32)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Int32Type},
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
	assert.Equal(t, field{int16(2), key, Int16Type}, f)
}

func TestInt16p(t *testing.T) {
	v := int16(2)
	type args struct {
		value *int16
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*int16)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Int16Type},
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
	assert.Equal(t, field{int8(2), key, Int8Type}, f)
}

func TestInt8p(t *testing.T) {
	v := int8(2)
	type args struct {
		value *int8
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*int8)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Int8Type},
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
	assert.Equal(t, field{"val", key, StringType}, f)
}

func TestStringp(t *testing.T) {
	v := "val"
	type args struct {
		value *string
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*string)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", StringType},
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
	assert.Equal(t, field{val, key, TimeType}, f)
}

func TestTimep(t *testing.T) {
	v := time.Now()
	type args struct {
		value *time.Time
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*time.Time)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", TimeType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Timep(tt.args.key, tt.args.value), "Timep(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func TestNewContext(t *testing.T) {
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
	ctx = NewContext(ctx, sf("a", "b"))
	fields, ok = ctx.Value(fieldKey{}).([]Field)
	assert.True(t, ok)
	assert.Len(t, fields, 2)
	assert.Equal(t, field, fields[0])
	assert.Equal(t, sf("a", "b"), fields[1])
}

func TestUint(t *testing.T) {
	key := "key"
	f := Uint(key, uint(2))
	assert.Equal(t, field{uint64(2), key, Uint64Type}, f)
}

func TestUintp(t *testing.T) {
	v := uint(2)
	type args struct {
		value *uint
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uint64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{uint64(v), "key", Uint64Type},
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
	assert.Equal(t, field{uint64(2), key, Uint64Type}, f)
}

func TestUint64p(t *testing.T) {
	v := uint64(2)
	type args struct {
		value *uint64
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uint64)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Uint64Type},
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
	assert.Equal(t, field{uint32(2), key, Uint32Type}, f)
}

func TestUint32p(t *testing.T) {
	v := uint32(2)
	type args struct {
		value *uint32
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uint32)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Uint32Type},
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
	assert.Equal(t, field{uint16(2), key, Uint16Type}, f)
}

func TestUint16p(t *testing.T) {
	v := uint16(2)
	type args struct {
		value *uint16
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uint16)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Uint16Type},
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
	assert.Equal(t, field{uint8(2), key, Uint8Type}, f)
}

func TestUint8p(t *testing.T) {
	v := uint8(2)
	type args struct {
		value *uint8
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uint8)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", Uint8Type},
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
	assert.Equal(t, field{val, key, UintptrType}, f)
}

func TestUintptrp(t *testing.T) {
	s := "some"
	v := (uintptr)(unsafe.Pointer(&s))
	type args struct {
		value *uintptr
		key   string
	}
	tests := []struct {
		args args
		want Field
		name string
	}{
		{
			name: "nil",
			args: args{nil, "key"},
			want: field{(*uintptr)(nil), "key", UnknownType},
		},
		{
			name: "v",
			args: args{&v, "key"},
			want: field{v, "key", UintptrType},
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
	assert.Equal(t, field{val, key, StringerType}, f)
}

func TestError(t *testing.T) {
	val := io.EOF
	f := Error(val)
	assert.Equal(t, field{val, "error", ErrorType}, f)
}

func TestNamedError(t *testing.T) {
	key := "key"
	val := io.EOF
	f := NamedError(key, val)
	assert.Equal(t, field{val, key, ErrorType}, f)
}

func TestStack(t *testing.T) {
	key := "key"
	f := Stack(key)
	assert.Equal(t, field{0, "key", StackType}, f)
}

func TestStackSkip(t *testing.T) {
	key := "key"
	f := StackSkip(key, 3)
	assert.Equal(t, field{3, "key", StackType}, f)
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
