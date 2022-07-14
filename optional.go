// Code generated. DO NOT EDIT.

package optional

import (
	"time"
)

// Bool represents an optional bool value.
type Bool struct {
	value bool
	isSet bool
}

// MakeBool makes a new optional bool value set to the given value.
func MakeBool(value bool) Bool {
	return Bool{value, true}
}

// Get returns the bool value and a boolean value indicating whether the bool value is set.
func (b Bool) Get() (bool, bool) { return b.value, b.isSet }

// Set sets the bool value.
func (b *Bool) Set(value bool) {
	*b = MakeBool(value)
}

// Clear clears the bool value.
func (b *Bool) Clear() {
	*b = Bool{}
}

// IsSet returns a boolean value indicating whether the bool value is set.
func (b Bool) IsSet() bool { return b.isSet }

// Byte represents an optional byte value.
type Byte struct {
	value byte
	isSet bool
}

// MakeByte makes a new optional byte value set to the given value.
func MakeByte(value byte) Byte {
	return Byte{value, true}
}

// Get returns the byte value and a boolean value indicating whether the byte value is set.
func (b Byte) Get() (byte, bool) { return b.value, b.isSet }

// Set sets the byte value.
func (b *Byte) Set(value byte) {
	*b = MakeByte(value)
}

// Clear clears the byte value.
func (b *Byte) Clear() {
	*b = Byte{}
}

// IsSet returns a boolean value indicating whether the byte value is set.
func (b Byte) IsSet() bool { return b.isSet }

// Int represents an optional int value.
type Int struct {
	value int
	isSet bool
}

// MakeInt makes a new optional int value set to the given value.
func MakeInt(value int) Int {
	return Int{value, true}
}

// Get returns the int value and a boolean value indicating whether the int value is set.
func (i Int) Get() (int, bool) { return i.value, i.isSet }

// Set sets the int value.
func (i *Int) Set(value int) {
	*i = MakeInt(value)
}

// Clear clears the int value.
func (i *Int) Clear() {
	*i = Int{}
}

// IsSet returns a boolean value indicating whether the int value is set.
func (i Int) IsSet() bool { return i.isSet }

// Uint represents an optional uint value.
type Uint struct {
	value uint
	isSet bool
}

// MakeUint makes a new optional uint value set to the given value.
func MakeUint(value uint) Uint {
	return Uint{value, true}
}

// Get returns the uint value and a boolean value indicating whether the uint value is set.
func (u Uint) Get() (uint, bool) { return u.value, u.isSet }

// Set sets the uint value.
func (u *Uint) Set(value uint) {
	*u = MakeUint(value)
}

// Clear clears the uint value.
func (u *Uint) Clear() {
	*u = Uint{}
}

// IsSet returns a boolean value indicating whether the uint value is set.
func (u Uint) IsSet() bool { return u.isSet }

// Int8 represents an optional int8 value.
type Int8 struct {
	value int8
	isSet bool
}

// MakeInt8 makes a new optional int8 value set to the given value.
func MakeInt8(value int8) Int8 {
	return Int8{value, true}
}

// Get returns the int8 value and a boolean value indicating whether the int8 value is set.
func (i Int8) Get() (int8, bool) { return i.value, i.isSet }

// Set sets the int8 value.
func (i *Int8) Set(value int8) {
	*i = MakeInt8(value)
}

// Clear clears the int8 value.
func (i *Int8) Clear() {
	*i = Int8{}
}

// IsSet returns a boolean value indicating whether the int8 value is set.
func (i Int8) IsSet() bool { return i.isSet }

// Uint8 represents an optional uint8 value.
type Uint8 struct {
	value uint8
	isSet bool
}

// MakeUint8 makes a new optional uint8 value set to the given value.
func MakeUint8(value uint8) Uint8 {
	return Uint8{value, true}
}

// Get returns the uint8 value and a boolean value indicating whether the uint8 value is set.
func (u Uint8) Get() (uint8, bool) { return u.value, u.isSet }

// Set sets the uint8 value.
func (u *Uint8) Set(value uint8) {
	*u = MakeUint8(value)
}

// Clear clears the uint8 value.
func (u *Uint8) Clear() {
	*u = Uint8{}
}

// IsSet returns a boolean value indicating whether the uint8 value is set.
func (u Uint8) IsSet() bool { return u.isSet }

// Int16 represents an optional int16 value.
type Int16 struct {
	value int16
	isSet bool
}

// MakeInt16 makes a new optional int16 value set to the given value.
func MakeInt16(value int16) Int16 {
	return Int16{value, true}
}

// Get returns the int16 value and a boolean value indicating whether the int16 value is set.
func (i Int16) Get() (int16, bool) { return i.value, i.isSet }

// Set sets the int16 value.
func (i *Int16) Set(value int16) {
	*i = MakeInt16(value)
}

// Clear clears the int16 value.
func (i *Int16) Clear() {
	*i = Int16{}
}

// IsSet returns a boolean value indicating whether the int16 value is set.
func (i Int16) IsSet() bool { return i.isSet }

// Uint16 represents an optional uint16 value.
type Uint16 struct {
	value uint16
	isSet bool
}

// MakeUint16 makes a new optional uint16 value set to the given value.
func MakeUint16(value uint16) Uint16 {
	return Uint16{value, true}
}

// Get returns the uint16 value and a boolean value indicating whether the uint16 value is set.
func (u Uint16) Get() (uint16, bool) { return u.value, u.isSet }

// Set sets the uint16 value.
func (u *Uint16) Set(value uint16) {
	*u = MakeUint16(value)
}

// Clear clears the uint16 value.
func (u *Uint16) Clear() {
	*u = Uint16{}
}

// IsSet returns a boolean value indicating whether the uint16 value is set.
func (u Uint16) IsSet() bool { return u.isSet }

// Int32 represents an optional int32 value.
type Int32 struct {
	value int32
	isSet bool
}

// MakeInt32 makes a new optional int32 value set to the given value.
func MakeInt32(value int32) Int32 {
	return Int32{value, true}
}

// Get returns the int32 value and a boolean value indicating whether the int32 value is set.
func (i Int32) Get() (int32, bool) { return i.value, i.isSet }

// Set sets the int32 value.
func (i *Int32) Set(value int32) {
	*i = MakeInt32(value)
}

// Clear clears the int32 value.
func (i *Int32) Clear() {
	*i = Int32{}
}

// IsSet returns a boolean value indicating whether the int32 value is set.
func (i Int32) IsSet() bool { return i.isSet }

// Uint32 represents an optional uint32 value.
type Uint32 struct {
	value uint32
	isSet bool
}

// MakeUint32 makes a new optional uint32 value set to the given value.
func MakeUint32(value uint32) Uint32 {
	return Uint32{value, true}
}

// Get returns the uint32 value and a boolean value indicating whether the uint32 value is set.
func (u Uint32) Get() (uint32, bool) { return u.value, u.isSet }

// Set sets the uint32 value.
func (u *Uint32) Set(value uint32) {
	*u = MakeUint32(value)
}

// Clear clears the uint32 value.
func (u *Uint32) Clear() {
	*u = Uint32{}
}

// IsSet returns a boolean value indicating whether the uint32 value is set.
func (u Uint32) IsSet() bool { return u.isSet }

// Int64 represents an optional int64 value.
type Int64 struct {
	value int64
	isSet bool
}

// MakeInt64 makes a new optional int64 value set to the given value.
func MakeInt64(value int64) Int64 {
	return Int64{value, true}
}

// Get returns the int64 value and a boolean value indicating whether the int64 value is set.
func (i Int64) Get() (int64, bool) { return i.value, i.isSet }

// Set sets the int64 value.
func (i *Int64) Set(value int64) {
	*i = MakeInt64(value)
}

// Clear clears the int64 value.
func (i *Int64) Clear() {
	*i = Int64{}
}

// IsSet returns a boolean value indicating whether the int64 value is set.
func (i Int64) IsSet() bool { return i.isSet }

// Uint64 represents an optional uint64 value.
type Uint64 struct {
	value uint64
	isSet bool
}

// MakeUint64 makes a new optional uint64 value set to the given value.
func MakeUint64(value uint64) Uint64 {
	return Uint64{value, true}
}

// Get returns the uint64 value and a boolean value indicating whether the uint64 value is set.
func (u Uint64) Get() (uint64, bool) { return u.value, u.isSet }

// Set sets the uint64 value.
func (u *Uint64) Set(value uint64) {
	*u = MakeUint64(value)
}

// Clear clears the uint64 value.
func (u *Uint64) Clear() {
	*u = Uint64{}
}

// IsSet returns a boolean value indicating whether the uint64 value is set.
func (u Uint64) IsSet() bool { return u.isSet }

// Uintptr represents an optional uintptr value.
type Uintptr struct {
	value uintptr
	isSet bool
}

// MakeUintptr makes a new optional uintptr value set to the given value.
func MakeUintptr(value uintptr) Uintptr {
	return Uintptr{value, true}
}

// Get returns the uintptr value and a boolean value indicating whether the uintptr value is set.
func (u Uintptr) Get() (uintptr, bool) { return u.value, u.isSet }

// Set sets the uintptr value.
func (u *Uintptr) Set(value uintptr) {
	*u = MakeUintptr(value)
}

// Clear clears the uintptr value.
func (u *Uintptr) Clear() {
	*u = Uintptr{}
}

// IsSet returns a boolean value indicating whether the uintptr value is set.
func (u Uintptr) IsSet() bool { return u.isSet }

// Float32 represents an optional float32 value.
type Float32 struct {
	value float32
	isSet bool
}

// MakeFloat32 makes a new optional float32 value set to the given value.
func MakeFloat32(value float32) Float32 {
	return Float32{value, true}
}

// Get returns the float32 value and a boolean value indicating whether the float32 value is set.
func (f Float32) Get() (float32, bool) { return f.value, f.isSet }

// Set sets the float32 value.
func (f *Float32) Set(value float32) {
	*f = MakeFloat32(value)
}

// Clear clears the float32 value.
func (f *Float32) Clear() {
	*f = Float32{}
}

// IsSet returns a boolean value indicating whether the float32 value is set.
func (f Float32) IsSet() bool { return f.isSet }

// Float64 represents an optional float64 value.
type Float64 struct {
	value float64
	isSet bool
}

// MakeFloat64 makes a new optional float64 value set to the given value.
func MakeFloat64(value float64) Float64 {
	return Float64{value, true}
}

// Get returns the float64 value and a boolean value indicating whether the float64 value is set.
func (f Float64) Get() (float64, bool) { return f.value, f.isSet }

// Set sets the float64 value.
func (f *Float64) Set(value float64) {
	*f = MakeFloat64(value)
}

// Clear clears the float64 value.
func (f *Float64) Clear() {
	*f = Float64{}
}

// IsSet returns a boolean value indicating whether the float64 value is set.
func (f Float64) IsSet() bool { return f.isSet }

// Complex64 represents an optional complex64 value.
type Complex64 struct {
	value complex64
	isSet bool
}

// MakeComplex64 makes a new optional complex64 value set to the given value.
func MakeComplex64(value complex64) Complex64 {
	return Complex64{value, true}
}

// Get returns the complex64 value and a boolean value indicating whether the complex64 value is set.
func (c Complex64) Get() (complex64, bool) { return c.value, c.isSet }

// Set sets the complex64 value.
func (c *Complex64) Set(value complex64) {
	*c = MakeComplex64(value)
}

// Clear clears the complex64 value.
func (c *Complex64) Clear() {
	*c = Complex64{}
}

// IsSet returns a boolean value indicating whether the complex64 value is set.
func (c Complex64) IsSet() bool { return c.isSet }

// Complex128 represents an optional complex128 value.
type Complex128 struct {
	value complex128
	isSet bool
}

// MakeComplex128 makes a new optional complex128 value set to the given value.
func MakeComplex128(value complex128) Complex128 {
	return Complex128{value, true}
}

// Get returns the complex128 value and a boolean value indicating whether the complex128 value is set.
func (c Complex128) Get() (complex128, bool) { return c.value, c.isSet }

// Set sets the complex128 value.
func (c *Complex128) Set(value complex128) {
	*c = MakeComplex128(value)
}

// Clear clears the complex128 value.
func (c *Complex128) Clear() {
	*c = Complex128{}
}

// IsSet returns a boolean value indicating whether the complex128 value is set.
func (c Complex128) IsSet() bool { return c.isSet }

// Rune represents an optional rune value.
type Rune struct {
	value rune
	isSet bool
}

// MakeRune makes a new optional rune value set to the given value.
func MakeRune(value rune) Rune {
	return Rune{value, true}
}

// Get returns the rune value and a boolean value indicating whether the rune value is set.
func (r Rune) Get() (rune, bool) { return r.value, r.isSet }

// Set sets the rune value.
func (r *Rune) Set(value rune) {
	*r = MakeRune(value)
}

// Clear clears the rune value.
func (r *Rune) Clear() {
	*r = Rune{}
}

// IsSet returns a boolean value indicating whether the rune value is set.
func (r Rune) IsSet() bool { return r.isSet }

// String represents an optional string value.
type String struct {
	value string
	isSet bool
}

// MakeString makes a new optional string value set to the given value.
func MakeString(value string) String {
	return String{value, true}
}

// Get returns the string value and a boolean value indicating whether the string value is set.
func (s String) Get() (string, bool) { return s.value, s.isSet }

// Set sets the string value.
func (s *String) Set(value string) {
	*s = MakeString(value)
}

// Clear clears the string value.
func (s *String) Clear() {
	*s = String{}
}

// IsSet returns a boolean value indicating whether the string value is set.
func (s String) IsSet() bool { return s.isSet }

// Duration represents an optional time.Duration value.
type Duration struct {
	value time.Duration
	isSet bool
}

// MakeDuration makes a new optional time.Duration value set to the given value.
func MakeDuration(value time.Duration) Duration {
	return Duration{value, true}
}

// Get returns the time.Duration value and a boolean value indicating whether the time.Duration value is set.
func (d Duration) Get() (time.Duration, bool) { return d.value, d.isSet }

// Set sets the time.Duration value.
func (d *Duration) Set(value time.Duration) {
	*d = MakeDuration(value)
}

// Clear clears the time.Duration value.
func (d *Duration) Clear() {
	*d = Duration{}
}

// IsSet returns a boolean value indicating whether the time.Duration value is set.
func (d Duration) IsSet() bool { return d.isSet }
