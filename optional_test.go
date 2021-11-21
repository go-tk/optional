// Code generated. DO NOT EDIT.

package optional_test

import (
	"testing"
	"time"

	. "github.com/go-tk/optional"
)

func TestBool_Set(t *testing.T) {
	var b Bool
	b.Set(true)
	if b != MakeBool(true) {
		t.Error("should be expected value")
	}
}

func TestBool_Clear(t *testing.T) {
	b := MakeBool(true)
	b.Clear()
	if b != (Bool{}) {
		t.Error("should be zero value")
	}
}

func TestBool_Value(t *testing.T) {
	{
		b := MakeBool(true)
		if b.Value() != true {
			t.Error("should be expected value")
		}
	}
	{
		var b Bool
		if b.Value() != false {
			t.Error("should be zero value")
		}
	}
	{
		var b Bool
		b.Set(true)
		if b.Value() != true {
			t.Error("should be expected value")
		}
	}
	{
		b := MakeBool(true)
		b.Clear()
		if b.Value() != false {
			t.Error("should be zero value")
		}
	}
}

func TestBool_HasValue(t *testing.T) {
	{
		b := MakeBool(true)
		if !b.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var b Bool
		if b.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var b Bool
		b.Set(true)
		if !b.HasValue() {
			t.Error("should have value")
		}
	}
	{
		b := MakeBool(true)
		b.Clear()
		if b.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestByte_Set(t *testing.T) {
	var b Byte
	b.Set('b')
	if b != MakeByte('b') {
		t.Error("should be expected value")
	}
}

func TestByte_Clear(t *testing.T) {
	b := MakeByte('b')
	b.Clear()
	if b != (Byte{}) {
		t.Error("should be zero value")
	}
}

func TestByte_Value(t *testing.T) {
	{
		b := MakeByte('b')
		if b.Value() != 'b' {
			t.Error("should be expected value")
		}
	}
	{
		var b Byte
		if b.Value() != '\000' {
			t.Error("should be zero value")
		}
	}
	{
		var b Byte
		b.Set('b')
		if b.Value() != 'b' {
			t.Error("should be expected value")
		}
	}
	{
		b := MakeByte('b')
		b.Clear()
		if b.Value() != '\000' {
			t.Error("should be zero value")
		}
	}
}

func TestByte_HasValue(t *testing.T) {
	{
		b := MakeByte('b')
		if !b.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var b Byte
		if b.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var b Byte
		b.Set('b')
		if !b.HasValue() {
			t.Error("should have value")
		}
	}
	{
		b := MakeByte('b')
		b.Clear()
		if b.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestInt_Set(t *testing.T) {
	var i Int
	i.Set(100)
	if i != MakeInt(100) {
		t.Error("should be expected value")
	}
}

func TestInt_Clear(t *testing.T) {
	i := MakeInt(100)
	i.Clear()
	if i != (Int{}) {
		t.Error("should be zero value")
	}
}

func TestInt_Value(t *testing.T) {
	{
		i := MakeInt(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var i Int
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var i Int
		i.Set(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		i := MakeInt(100)
		i.Clear()
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestInt_HasValue(t *testing.T) {
	{
		i := MakeInt(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var i Int
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var i Int
		i.Set(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		i := MakeInt(100)
		i.Clear()
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUint_Set(t *testing.T) {
	var u Uint
	u.Set(100)
	if u != MakeUint(100) {
		t.Error("should be expected value")
	}
}

func TestUint_Clear(t *testing.T) {
	u := MakeUint(100)
	u.Clear()
	if u != (Uint{}) {
		t.Error("should be zero value")
	}
}

func TestUint_Value(t *testing.T) {
	{
		u := MakeUint(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uint
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uint
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUint(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUint_HasValue(t *testing.T) {
	{
		u := MakeUint(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uint
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uint
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUint(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestInt8_Set(t *testing.T) {
	var i Int8
	i.Set(100)
	if i != MakeInt8(100) {
		t.Error("should be expected value")
	}
}

func TestInt8_Clear(t *testing.T) {
	i := MakeInt8(100)
	i.Clear()
	if i != (Int8{}) {
		t.Error("should be zero value")
	}
}

func TestInt8_Value(t *testing.T) {
	{
		i := MakeInt8(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var i Int8
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var i Int8
		i.Set(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		i := MakeInt8(100)
		i.Clear()
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestInt8_HasValue(t *testing.T) {
	{
		i := MakeInt8(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var i Int8
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var i Int8
		i.Set(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		i := MakeInt8(100)
		i.Clear()
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUint8_Set(t *testing.T) {
	var u Uint8
	u.Set(100)
	if u != MakeUint8(100) {
		t.Error("should be expected value")
	}
}

func TestUint8_Clear(t *testing.T) {
	u := MakeUint8(100)
	u.Clear()
	if u != (Uint8{}) {
		t.Error("should be zero value")
	}
}

func TestUint8_Value(t *testing.T) {
	{
		u := MakeUint8(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uint8
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uint8
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUint8(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUint8_HasValue(t *testing.T) {
	{
		u := MakeUint8(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uint8
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uint8
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUint8(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestInt16_Set(t *testing.T) {
	var i Int16
	i.Set(100)
	if i != MakeInt16(100) {
		t.Error("should be expected value")
	}
}

func TestInt16_Clear(t *testing.T) {
	i := MakeInt16(100)
	i.Clear()
	if i != (Int16{}) {
		t.Error("should be zero value")
	}
}

func TestInt16_Value(t *testing.T) {
	{
		i := MakeInt16(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var i Int16
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var i Int16
		i.Set(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		i := MakeInt16(100)
		i.Clear()
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestInt16_HasValue(t *testing.T) {
	{
		i := MakeInt16(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var i Int16
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var i Int16
		i.Set(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		i := MakeInt16(100)
		i.Clear()
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUint16_Set(t *testing.T) {
	var u Uint16
	u.Set(100)
	if u != MakeUint16(100) {
		t.Error("should be expected value")
	}
}

func TestUint16_Clear(t *testing.T) {
	u := MakeUint16(100)
	u.Clear()
	if u != (Uint16{}) {
		t.Error("should be zero value")
	}
}

func TestUint16_Value(t *testing.T) {
	{
		u := MakeUint16(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uint16
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uint16
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUint16(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUint16_HasValue(t *testing.T) {
	{
		u := MakeUint16(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uint16
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uint16
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUint16(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestInt32_Set(t *testing.T) {
	var i Int32
	i.Set(100)
	if i != MakeInt32(100) {
		t.Error("should be expected value")
	}
}

func TestInt32_Clear(t *testing.T) {
	i := MakeInt32(100)
	i.Clear()
	if i != (Int32{}) {
		t.Error("should be zero value")
	}
}

func TestInt32_Value(t *testing.T) {
	{
		i := MakeInt32(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var i Int32
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var i Int32
		i.Set(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		i := MakeInt32(100)
		i.Clear()
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestInt32_HasValue(t *testing.T) {
	{
		i := MakeInt32(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var i Int32
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var i Int32
		i.Set(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		i := MakeInt32(100)
		i.Clear()
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUint32_Set(t *testing.T) {
	var u Uint32
	u.Set(100)
	if u != MakeUint32(100) {
		t.Error("should be expected value")
	}
}

func TestUint32_Clear(t *testing.T) {
	u := MakeUint32(100)
	u.Clear()
	if u != (Uint32{}) {
		t.Error("should be zero value")
	}
}

func TestUint32_Value(t *testing.T) {
	{
		u := MakeUint32(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uint32
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uint32
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUint32(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUint32_HasValue(t *testing.T) {
	{
		u := MakeUint32(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uint32
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uint32
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUint32(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestInt64_Set(t *testing.T) {
	var i Int64
	i.Set(100)
	if i != MakeInt64(100) {
		t.Error("should be expected value")
	}
}

func TestInt64_Clear(t *testing.T) {
	i := MakeInt64(100)
	i.Clear()
	if i != (Int64{}) {
		t.Error("should be zero value")
	}
}

func TestInt64_Value(t *testing.T) {
	{
		i := MakeInt64(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var i Int64
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var i Int64
		i.Set(100)
		if i.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		i := MakeInt64(100)
		i.Clear()
		if i.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestInt64_HasValue(t *testing.T) {
	{
		i := MakeInt64(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var i Int64
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var i Int64
		i.Set(100)
		if !i.HasValue() {
			t.Error("should have value")
		}
	}
	{
		i := MakeInt64(100)
		i.Clear()
		if i.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUint64_Set(t *testing.T) {
	var u Uint64
	u.Set(100)
	if u != MakeUint64(100) {
		t.Error("should be expected value")
	}
}

func TestUint64_Clear(t *testing.T) {
	u := MakeUint64(100)
	u.Clear()
	if u != (Uint64{}) {
		t.Error("should be zero value")
	}
}

func TestUint64_Value(t *testing.T) {
	{
		u := MakeUint64(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uint64
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uint64
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUint64(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUint64_HasValue(t *testing.T) {
	{
		u := MakeUint64(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uint64
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uint64
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUint64(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestUintptr_Set(t *testing.T) {
	var u Uintptr
	u.Set(100)
	if u != MakeUintptr(100) {
		t.Error("should be expected value")
	}
}

func TestUintptr_Clear(t *testing.T) {
	u := MakeUintptr(100)
	u.Clear()
	if u != (Uintptr{}) {
		t.Error("should be zero value")
	}
}

func TestUintptr_Value(t *testing.T) {
	{
		u := MakeUintptr(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		var u Uintptr
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var u Uintptr
		u.Set(100)
		if u.Value() != 100 {
			t.Error("should be expected value")
		}
	}
	{
		u := MakeUintptr(100)
		u.Clear()
		if u.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestUintptr_HasValue(t *testing.T) {
	{
		u := MakeUintptr(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var u Uintptr
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var u Uintptr
		u.Set(100)
		if !u.HasValue() {
			t.Error("should have value")
		}
	}
	{
		u := MakeUintptr(100)
		u.Clear()
		if u.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestFloat32_Set(t *testing.T) {
	var f Float32
	f.Set(1.0)
	if f != MakeFloat32(1.0) {
		t.Error("should be expected value")
	}
}

func TestFloat32_Clear(t *testing.T) {
	f := MakeFloat32(1.0)
	f.Clear()
	if f != (Float32{}) {
		t.Error("should be zero value")
	}
}

func TestFloat32_Value(t *testing.T) {
	{
		f := MakeFloat32(1.0)
		if f.Value() != 1.0 {
			t.Error("should be expected value")
		}
	}
	{
		var f Float32
		if f.Value() != 0.0 {
			t.Error("should be zero value")
		}
	}
	{
		var f Float32
		f.Set(1.0)
		if f.Value() != 1.0 {
			t.Error("should be expected value")
		}
	}
	{
		f := MakeFloat32(1.0)
		f.Clear()
		if f.Value() != 0.0 {
			t.Error("should be zero value")
		}
	}
}

func TestFloat32_HasValue(t *testing.T) {
	{
		f := MakeFloat32(1.0)
		if !f.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var f Float32
		if f.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var f Float32
		f.Set(1.0)
		if !f.HasValue() {
			t.Error("should have value")
		}
	}
	{
		f := MakeFloat32(1.0)
		f.Clear()
		if f.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestFloat64_Set(t *testing.T) {
	var f Float64
	f.Set(1.0)
	if f != MakeFloat64(1.0) {
		t.Error("should be expected value")
	}
}

func TestFloat64_Clear(t *testing.T) {
	f := MakeFloat64(1.0)
	f.Clear()
	if f != (Float64{}) {
		t.Error("should be zero value")
	}
}

func TestFloat64_Value(t *testing.T) {
	{
		f := MakeFloat64(1.0)
		if f.Value() != 1.0 {
			t.Error("should be expected value")
		}
	}
	{
		var f Float64
		if f.Value() != 0.0 {
			t.Error("should be zero value")
		}
	}
	{
		var f Float64
		f.Set(1.0)
		if f.Value() != 1.0 {
			t.Error("should be expected value")
		}
	}
	{
		f := MakeFloat64(1.0)
		f.Clear()
		if f.Value() != 0.0 {
			t.Error("should be zero value")
		}
	}
}

func TestFloat64_HasValue(t *testing.T) {
	{
		f := MakeFloat64(1.0)
		if !f.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var f Float64
		if f.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var f Float64
		f.Set(1.0)
		if !f.HasValue() {
			t.Error("should have value")
		}
	}
	{
		f := MakeFloat64(1.0)
		f.Clear()
		if f.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestComplex64_Set(t *testing.T) {
	var c Complex64
	c.Set(complex(1, 2))
	if c != MakeComplex64(complex(1, 2)) {
		t.Error("should be expected value")
	}
}

func TestComplex64_Clear(t *testing.T) {
	c := MakeComplex64(complex(1, 2))
	c.Clear()
	if c != (Complex64{}) {
		t.Error("should be zero value")
	}
}

func TestComplex64_Value(t *testing.T) {
	{
		c := MakeComplex64(complex(1, 2))
		if c.Value() != complex(1, 2) {
			t.Error("should be expected value")
		}
	}
	{
		var c Complex64
		if c.Value() != complex(0, 0) {
			t.Error("should be zero value")
		}
	}
	{
		var c Complex64
		c.Set(complex(1, 2))
		if c.Value() != complex(1, 2) {
			t.Error("should be expected value")
		}
	}
	{
		c := MakeComplex64(complex(1, 2))
		c.Clear()
		if c.Value() != complex(0, 0) {
			t.Error("should be zero value")
		}
	}
}

func TestComplex64_HasValue(t *testing.T) {
	{
		c := MakeComplex64(complex(1, 2))
		if !c.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var c Complex64
		if c.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var c Complex64
		c.Set(complex(1, 2))
		if !c.HasValue() {
			t.Error("should have value")
		}
	}
	{
		c := MakeComplex64(complex(1, 2))
		c.Clear()
		if c.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestComplex128_Set(t *testing.T) {
	var c Complex128
	c.Set(complex(1, 2))
	if c != MakeComplex128(complex(1, 2)) {
		t.Error("should be expected value")
	}
}

func TestComplex128_Clear(t *testing.T) {
	c := MakeComplex128(complex(1, 2))
	c.Clear()
	if c != (Complex128{}) {
		t.Error("should be zero value")
	}
}

func TestComplex128_Value(t *testing.T) {
	{
		c := MakeComplex128(complex(1, 2))
		if c.Value() != complex(1, 2) {
			t.Error("should be expected value")
		}
	}
	{
		var c Complex128
		if c.Value() != complex(0, 0) {
			t.Error("should be zero value")
		}
	}
	{
		var c Complex128
		c.Set(complex(1, 2))
		if c.Value() != complex(1, 2) {
			t.Error("should be expected value")
		}
	}
	{
		c := MakeComplex128(complex(1, 2))
		c.Clear()
		if c.Value() != complex(0, 0) {
			t.Error("should be zero value")
		}
	}
}

func TestComplex128_HasValue(t *testing.T) {
	{
		c := MakeComplex128(complex(1, 2))
		if !c.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var c Complex128
		if c.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var c Complex128
		c.Set(complex(1, 2))
		if !c.HasValue() {
			t.Error("should have value")
		}
	}
	{
		c := MakeComplex128(complex(1, 2))
		c.Clear()
		if c.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestRune_Set(t *testing.T) {
	var r Rune
	r.Set('😊')
	if r != MakeRune('😊') {
		t.Error("should be expected value")
	}
}

func TestRune_Clear(t *testing.T) {
	r := MakeRune('😊')
	r.Clear()
	if r != (Rune{}) {
		t.Error("should be zero value")
	}
}

func TestRune_Value(t *testing.T) {
	{
		r := MakeRune('😊')
		if r.Value() != '😊' {
			t.Error("should be expected value")
		}
	}
	{
		var r Rune
		if r.Value() != '\000' {
			t.Error("should be zero value")
		}
	}
	{
		var r Rune
		r.Set('😊')
		if r.Value() != '😊' {
			t.Error("should be expected value")
		}
	}
	{
		r := MakeRune('😊')
		r.Clear()
		if r.Value() != '\000' {
			t.Error("should be zero value")
		}
	}
}

func TestRune_HasValue(t *testing.T) {
	{
		r := MakeRune('😊')
		if !r.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var r Rune
		if r.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var r Rune
		r.Set('😊')
		if !r.HasValue() {
			t.Error("should have value")
		}
	}
	{
		r := MakeRune('😊')
		r.Clear()
		if r.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestString_Set(t *testing.T) {
	var s String
	s.Set("foo")
	if s != MakeString("foo") {
		t.Error("should be expected value")
	}
}

func TestString_Clear(t *testing.T) {
	s := MakeString("foo")
	s.Clear()
	if s != (String{}) {
		t.Error("should be zero value")
	}
}

func TestString_Value(t *testing.T) {
	{
		s := MakeString("foo")
		if s.Value() != "foo" {
			t.Error("should be expected value")
		}
	}
	{
		var s String
		if s.Value() != "" {
			t.Error("should be zero value")
		}
	}
	{
		var s String
		s.Set("foo")
		if s.Value() != "foo" {
			t.Error("should be expected value")
		}
	}
	{
		s := MakeString("foo")
		s.Clear()
		if s.Value() != "" {
			t.Error("should be zero value")
		}
	}
}

func TestString_HasValue(t *testing.T) {
	{
		s := MakeString("foo")
		if !s.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var s String
		if s.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var s String
		s.Set("foo")
		if !s.HasValue() {
			t.Error("should have value")
		}
	}
	{
		s := MakeString("foo")
		s.Clear()
		if s.HasValue() {
			t.Error("should have no value")
		}
	}
}

func TestDuration_Set(t *testing.T) {
	var d Duration
	d.Set(100 * time.Second)
	if d != MakeDuration(100*time.Second) {
		t.Error("should be expected value")
	}
}

func TestDuration_Clear(t *testing.T) {
	d := MakeDuration(100 * time.Second)
	d.Clear()
	if d != (Duration{}) {
		t.Error("should be zero value")
	}
}

func TestDuration_Value(t *testing.T) {
	{
		d := MakeDuration(100 * time.Second)
		if d.Value() != 100*time.Second {
			t.Error("should be expected value")
		}
	}
	{
		var d Duration
		if d.Value() != 0 {
			t.Error("should be zero value")
		}
	}
	{
		var d Duration
		d.Set(100 * time.Second)
		if d.Value() != 100*time.Second {
			t.Error("should be expected value")
		}
	}
	{
		d := MakeDuration(100 * time.Second)
		d.Clear()
		if d.Value() != 0 {
			t.Error("should be zero value")
		}
	}
}

func TestDuration_HasValue(t *testing.T) {
	{
		d := MakeDuration(100 * time.Second)
		if !d.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var d Duration
		if d.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var d Duration
		d.Set(100 * time.Second)
		if !d.HasValue() {
			t.Error("should have value")
		}
	}
	{
		d := MakeDuration(100 * time.Second)
		d.Clear()
		if d.HasValue() {
			t.Error("should have no value")
		}
	}
}
