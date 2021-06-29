// Code generated. DO NOT EDIT.

package optional

type Bool struct {
	value    bool
	hasValue bool
}

func MakeBool(value bool) Bool {
	return Bool{value, true}
}

func (b Bool) Value() bool    { return b.value }
func (b Bool) HasValue() bool { return b.hasValue }

type Byte struct {
	value    byte
	hasValue bool
}

func MakeByte(value byte) Byte {
	return Byte{value, true}
}

func (b Byte) Value() byte    { return b.value }
func (b Byte) HasValue() bool { return b.hasValue }

type Int struct {
	value    int
	hasValue bool
}

func MakeInt(value int) Int {
	return Int{value, true}
}

func (i Int) Value() int     { return i.value }
func (i Int) HasValue() bool { return i.hasValue }

type Uint struct {
	value    uint
	hasValue bool
}

func MakeUint(value uint) Uint {
	return Uint{value, true}
}

func (u Uint) Value() uint    { return u.value }
func (u Uint) HasValue() bool { return u.hasValue }

type Int8 struct {
	value    int8
	hasValue bool
}

func MakeInt8(value int8) Int8 {
	return Int8{value, true}
}

func (i Int8) Value() int8    { return i.value }
func (i Int8) HasValue() bool { return i.hasValue }

type Uint8 struct {
	value    uint8
	hasValue bool
}

func MakeUint8(value uint8) Uint8 {
	return Uint8{value, true}
}

func (u Uint8) Value() uint8   { return u.value }
func (u Uint8) HasValue() bool { return u.hasValue }

type Int16 struct {
	value    int16
	hasValue bool
}

func MakeInt16(value int16) Int16 {
	return Int16{value, true}
}

func (i Int16) Value() int16   { return i.value }
func (i Int16) HasValue() bool { return i.hasValue }

type Uint16 struct {
	value    uint16
	hasValue bool
}

func MakeUint16(value uint16) Uint16 {
	return Uint16{value, true}
}

func (u Uint16) Value() uint16  { return u.value }
func (u Uint16) HasValue() bool { return u.hasValue }

type Int32 struct {
	value    int32
	hasValue bool
}

func MakeInt32(value int32) Int32 {
	return Int32{value, true}
}

func (i Int32) Value() int32   { return i.value }
func (i Int32) HasValue() bool { return i.hasValue }

type Uint32 struct {
	value    uint32
	hasValue bool
}

func MakeUint32(value uint32) Uint32 {
	return Uint32{value, true}
}

func (u Uint32) Value() uint32  { return u.value }
func (u Uint32) HasValue() bool { return u.hasValue }

type Int64 struct {
	value    int64
	hasValue bool
}

func MakeInt64(value int64) Int64 {
	return Int64{value, true}
}

func (i Int64) Value() int64   { return i.value }
func (i Int64) HasValue() bool { return i.hasValue }

type Uint64 struct {
	value    uint64
	hasValue bool
}

func MakeUint64(value uint64) Uint64 {
	return Uint64{value, true}
}

func (u Uint64) Value() uint64  { return u.value }
func (u Uint64) HasValue() bool { return u.hasValue }

type Uintptr struct {
	value    uintptr
	hasValue bool
}

func MakeUintptr(value uintptr) Uintptr {
	return Uintptr{value, true}
}

func (u Uintptr) Value() uintptr { return u.value }
func (u Uintptr) HasValue() bool { return u.hasValue }

type Float32 struct {
	value    float32
	hasValue bool
}

func MakeFloat32(value float32) Float32 {
	return Float32{value, true}
}

func (f Float32) Value() float32 { return f.value }
func (f Float32) HasValue() bool { return f.hasValue }

type Float64 struct {
	value    float64
	hasValue bool
}

func MakeFloat64(value float64) Float64 {
	return Float64{value, true}
}

func (f Float64) Value() float64 { return f.value }
func (f Float64) HasValue() bool { return f.hasValue }

type Complex64 struct {
	value    complex64
	hasValue bool
}

func MakeComplex64(value complex64) Complex64 {
	return Complex64{value, true}
}

func (c Complex64) Value() complex64 { return c.value }
func (c Complex64) HasValue() bool   { return c.hasValue }

type Complex128 struct {
	value    complex128
	hasValue bool
}

func MakeComplex128(value complex128) Complex128 {
	return Complex128{value, true}
}

func (c Complex128) Value() complex128 { return c.value }
func (c Complex128) HasValue() bool    { return c.hasValue }

type Rune struct {
	value    rune
	hasValue bool
}

func MakeRune(value rune) Rune {
	return Rune{value, true}
}

func (r Rune) Value() rune    { return r.value }
func (r Rune) HasValue() bool { return r.hasValue }

type String struct {
	value    string
	hasValue bool
}

func MakeString(value string) String {
	return String{value, true}
}

func (s String) Value() string  { return s.value }
func (s String) HasValue() bool { return s.hasValue }
