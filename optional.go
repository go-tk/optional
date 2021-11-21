// Code generated. DO NOT EDIT.

package optional

import (
	"time"
)

// Bool represents an optional bool value.
type Bool struct {
	value    bool
	hasValue bool
}

// MakeBool makes a new optional bool value with the given value.
func MakeBool(value bool) Bool {
	return Bool{value, true}
}

// Set sets the bool value.
func (b *Bool) Set(value bool) {
	*b = MakeBool(value)
}

// Clear clears the bool value.
func (b *Bool) Clear() {
	*b = Bool{}
}

// Value returns the bool value.
func (b Bool) Value() bool { return b.value }

// HasValue returns a boolean value indicating whether the bool value is set.
func (b Bool) HasValue() bool { return b.hasValue }

// Byte represents an optional byte value.
type Byte struct {
	value    byte
	hasValue bool
}

// MakeByte makes a new optional byte value with the given value.
func MakeByte(value byte) Byte {
	return Byte{value, true}
}

// Set sets the byte value.
func (b *Byte) Set(value byte) {
	*b = MakeByte(value)
}

// Clear clears the byte value.
func (b *Byte) Clear() {
	*b = Byte{}
}

// Value returns the byte value.
func (b Byte) Value() byte { return b.value }

// HasValue returns a boolean value indicating whether the byte value is set.
func (b Byte) HasValue() bool { return b.hasValue }

// Int represents an optional int value.
type Int struct {
	value    int
	hasValue bool
}

// MakeInt makes a new optional int value with the given value.
func MakeInt(value int) Int {
	return Int{value, true}
}

// Set sets the int value.
func (i *Int) Set(value int) {
	*i = MakeInt(value)
}

// Clear clears the int value.
func (i *Int) Clear() {
	*i = Int{}
}

// Value returns the int value.
func (i Int) Value() int { return i.value }

// HasValue returns a boolean value indicating whether the int value is set.
func (i Int) HasValue() bool { return i.hasValue }

// Uint represents an optional uint value.
type Uint struct {
	value    uint
	hasValue bool
}

// MakeUint makes a new optional uint value with the given value.
func MakeUint(value uint) Uint {
	return Uint{value, true}
}

// Set sets the uint value.
func (u *Uint) Set(value uint) {
	*u = MakeUint(value)
}

// Clear clears the uint value.
func (u *Uint) Clear() {
	*u = Uint{}
}

// Value returns the uint value.
func (u Uint) Value() uint { return u.value }

// HasValue returns a boolean value indicating whether the uint value is set.
func (u Uint) HasValue() bool { return u.hasValue }

// Int8 represents an optional int8 value.
type Int8 struct {
	value    int8
	hasValue bool
}

// MakeInt8 makes a new optional int8 value with the given value.
func MakeInt8(value int8) Int8 {
	return Int8{value, true}
}

// Set sets the int8 value.
func (i *Int8) Set(value int8) {
	*i = MakeInt8(value)
}

// Clear clears the int8 value.
func (i *Int8) Clear() {
	*i = Int8{}
}

// Value returns the int8 value.
func (i Int8) Value() int8 { return i.value }

// HasValue returns a boolean value indicating whether the int8 value is set.
func (i Int8) HasValue() bool { return i.hasValue }

// Uint8 represents an optional uint8 value.
type Uint8 struct {
	value    uint8
	hasValue bool
}

// MakeUint8 makes a new optional uint8 value with the given value.
func MakeUint8(value uint8) Uint8 {
	return Uint8{value, true}
}

// Set sets the uint8 value.
func (u *Uint8) Set(value uint8) {
	*u = MakeUint8(value)
}

// Clear clears the uint8 value.
func (u *Uint8) Clear() {
	*u = Uint8{}
}

// Value returns the uint8 value.
func (u Uint8) Value() uint8 { return u.value }

// HasValue returns a boolean value indicating whether the uint8 value is set.
func (u Uint8) HasValue() bool { return u.hasValue }

// Int16 represents an optional int16 value.
type Int16 struct {
	value    int16
	hasValue bool
}

// MakeInt16 makes a new optional int16 value with the given value.
func MakeInt16(value int16) Int16 {
	return Int16{value, true}
}

// Set sets the int16 value.
func (i *Int16) Set(value int16) {
	*i = MakeInt16(value)
}

// Clear clears the int16 value.
func (i *Int16) Clear() {
	*i = Int16{}
}

// Value returns the int16 value.
func (i Int16) Value() int16 { return i.value }

// HasValue returns a boolean value indicating whether the int16 value is set.
func (i Int16) HasValue() bool { return i.hasValue }

// Uint16 represents an optional uint16 value.
type Uint16 struct {
	value    uint16
	hasValue bool
}

// MakeUint16 makes a new optional uint16 value with the given value.
func MakeUint16(value uint16) Uint16 {
	return Uint16{value, true}
}

// Set sets the uint16 value.
func (u *Uint16) Set(value uint16) {
	*u = MakeUint16(value)
}

// Clear clears the uint16 value.
func (u *Uint16) Clear() {
	*u = Uint16{}
}

// Value returns the uint16 value.
func (u Uint16) Value() uint16 { return u.value }

// HasValue returns a boolean value indicating whether the uint16 value is set.
func (u Uint16) HasValue() bool { return u.hasValue }

// Int32 represents an optional int32 value.
type Int32 struct {
	value    int32
	hasValue bool
}

// MakeInt32 makes a new optional int32 value with the given value.
func MakeInt32(value int32) Int32 {
	return Int32{value, true}
}

// Set sets the int32 value.
func (i *Int32) Set(value int32) {
	*i = MakeInt32(value)
}

// Clear clears the int32 value.
func (i *Int32) Clear() {
	*i = Int32{}
}

// Value returns the int32 value.
func (i Int32) Value() int32 { return i.value }

// HasValue returns a boolean value indicating whether the int32 value is set.
func (i Int32) HasValue() bool { return i.hasValue }

// Uint32 represents an optional uint32 value.
type Uint32 struct {
	value    uint32
	hasValue bool
}

// MakeUint32 makes a new optional uint32 value with the given value.
func MakeUint32(value uint32) Uint32 {
	return Uint32{value, true}
}

// Set sets the uint32 value.
func (u *Uint32) Set(value uint32) {
	*u = MakeUint32(value)
}

// Clear clears the uint32 value.
func (u *Uint32) Clear() {
	*u = Uint32{}
}

// Value returns the uint32 value.
func (u Uint32) Value() uint32 { return u.value }

// HasValue returns a boolean value indicating whether the uint32 value is set.
func (u Uint32) HasValue() bool { return u.hasValue }

// Int64 represents an optional int64 value.
type Int64 struct {
	value    int64
	hasValue bool
}

// MakeInt64 makes a new optional int64 value with the given value.
func MakeInt64(value int64) Int64 {
	return Int64{value, true}
}

// Set sets the int64 value.
func (i *Int64) Set(value int64) {
	*i = MakeInt64(value)
}

// Clear clears the int64 value.
func (i *Int64) Clear() {
	*i = Int64{}
}

// Value returns the int64 value.
func (i Int64) Value() int64 { return i.value }

// HasValue returns a boolean value indicating whether the int64 value is set.
func (i Int64) HasValue() bool { return i.hasValue }

// Uint64 represents an optional uint64 value.
type Uint64 struct {
	value    uint64
	hasValue bool
}

// MakeUint64 makes a new optional uint64 value with the given value.
func MakeUint64(value uint64) Uint64 {
	return Uint64{value, true}
}

// Set sets the uint64 value.
func (u *Uint64) Set(value uint64) {
	*u = MakeUint64(value)
}

// Clear clears the uint64 value.
func (u *Uint64) Clear() {
	*u = Uint64{}
}

// Value returns the uint64 value.
func (u Uint64) Value() uint64 { return u.value }

// HasValue returns a boolean value indicating whether the uint64 value is set.
func (u Uint64) HasValue() bool { return u.hasValue }

// Uintptr represents an optional uintptr value.
type Uintptr struct {
	value    uintptr
	hasValue bool
}

// MakeUintptr makes a new optional uintptr value with the given value.
func MakeUintptr(value uintptr) Uintptr {
	return Uintptr{value, true}
}

// Set sets the uintptr value.
func (u *Uintptr) Set(value uintptr) {
	*u = MakeUintptr(value)
}

// Clear clears the uintptr value.
func (u *Uintptr) Clear() {
	*u = Uintptr{}
}

// Value returns the uintptr value.
func (u Uintptr) Value() uintptr { return u.value }

// HasValue returns a boolean value indicating whether the uintptr value is set.
func (u Uintptr) HasValue() bool { return u.hasValue }

// Float32 represents an optional float32 value.
type Float32 struct {
	value    float32
	hasValue bool
}

// MakeFloat32 makes a new optional float32 value with the given value.
func MakeFloat32(value float32) Float32 {
	return Float32{value, true}
}

// Set sets the float32 value.
func (f *Float32) Set(value float32) {
	*f = MakeFloat32(value)
}

// Clear clears the float32 value.
func (f *Float32) Clear() {
	*f = Float32{}
}

// Value returns the float32 value.
func (f Float32) Value() float32 { return f.value }

// HasValue returns a boolean value indicating whether the float32 value is set.
func (f Float32) HasValue() bool { return f.hasValue }

// Float64 represents an optional float64 value.
type Float64 struct {
	value    float64
	hasValue bool
}

// MakeFloat64 makes a new optional float64 value with the given value.
func MakeFloat64(value float64) Float64 {
	return Float64{value, true}
}

// Set sets the float64 value.
func (f *Float64) Set(value float64) {
	*f = MakeFloat64(value)
}

// Clear clears the float64 value.
func (f *Float64) Clear() {
	*f = Float64{}
}

// Value returns the float64 value.
func (f Float64) Value() float64 { return f.value }

// HasValue returns a boolean value indicating whether the float64 value is set.
func (f Float64) HasValue() bool { return f.hasValue }

// Complex64 represents an optional complex64 value.
type Complex64 struct {
	value    complex64
	hasValue bool
}

// MakeComplex64 makes a new optional complex64 value with the given value.
func MakeComplex64(value complex64) Complex64 {
	return Complex64{value, true}
}

// Set sets the complex64 value.
func (c *Complex64) Set(value complex64) {
	*c = MakeComplex64(value)
}

// Clear clears the complex64 value.
func (c *Complex64) Clear() {
	*c = Complex64{}
}

// Value returns the complex64 value.
func (c Complex64) Value() complex64 { return c.value }

// HasValue returns a boolean value indicating whether the complex64 value is set.
func (c Complex64) HasValue() bool { return c.hasValue }

// Complex128 represents an optional complex128 value.
type Complex128 struct {
	value    complex128
	hasValue bool
}

// MakeComplex128 makes a new optional complex128 value with the given value.
func MakeComplex128(value complex128) Complex128 {
	return Complex128{value, true}
}

// Set sets the complex128 value.
func (c *Complex128) Set(value complex128) {
	*c = MakeComplex128(value)
}

// Clear clears the complex128 value.
func (c *Complex128) Clear() {
	*c = Complex128{}
}

// Value returns the complex128 value.
func (c Complex128) Value() complex128 { return c.value }

// HasValue returns a boolean value indicating whether the complex128 value is set.
func (c Complex128) HasValue() bool { return c.hasValue }

// Rune represents an optional rune value.
type Rune struct {
	value    rune
	hasValue bool
}

// MakeRune makes a new optional rune value with the given value.
func MakeRune(value rune) Rune {
	return Rune{value, true}
}

// Set sets the rune value.
func (r *Rune) Set(value rune) {
	*r = MakeRune(value)
}

// Clear clears the rune value.
func (r *Rune) Clear() {
	*r = Rune{}
}

// Value returns the rune value.
func (r Rune) Value() rune { return r.value }

// HasValue returns a boolean value indicating whether the rune value is set.
func (r Rune) HasValue() bool { return r.hasValue }

// String represents an optional string value.
type String struct {
	value    string
	hasValue bool
}

// MakeString makes a new optional string value with the given value.
func MakeString(value string) String {
	return String{value, true}
}

// Set sets the string value.
func (s *String) Set(value string) {
	*s = MakeString(value)
}

// Clear clears the string value.
func (s *String) Clear() {
	*s = String{}
}

// Value returns the string value.
func (s String) Value() string { return s.value }

// HasValue returns a boolean value indicating whether the string value is set.
func (s String) HasValue() bool { return s.hasValue }

// Duration represents an optional time.Duration value.
type Duration struct {
	value    time.Duration
	hasValue bool
}

// MakeDuration makes a new optional time.Duration value with the given value.
func MakeDuration(value time.Duration) Duration {
	return Duration{value, true}
}

// Set sets the time.Duration value.
func (d *Duration) Set(value time.Duration) {
	*d = MakeDuration(value)
}

// Clear clears the time.Duration value.
func (d *Duration) Clear() {
	*d = Duration{}
}

// Value returns the time.Duration value.
func (d Duration) Value() time.Duration { return d.value }

// HasValue returns a boolean value indicating whether the time.Duration value is set.
func (d Duration) HasValue() bool { return d.hasValue }
