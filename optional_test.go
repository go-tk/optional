// Code generated. DO NOT EDIT.

package optional_test

import (
	"testing"
	"time"

	. "github.com/go-tk/optional"
)

func TestBool_Get(t *testing.T) {
	{
		b := MakeBool(true)
		v, ok := b.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != true {
			t.Error("value should be expected")
		}
	}
	{
		var b Bool
		v, ok := b.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != false {
			t.Error("value should be zero")
		}
	}
	{
		var b Bool
		b.Set(true)
		v, ok := b.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != true {
			t.Error("value should be expected")
		}
	}
	{
		b := MakeBool(true)
		b.Clear()
		v, ok := b.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != false {
			t.Error("value should be zero")
		}
	}
}

func TestBool_Set(t *testing.T) {
	var b Bool
	b.Set(true)
	if b != MakeBool(true) {
		t.Error("value should be expected")
	}
}

func TestBool_Clear(t *testing.T) {
	b := MakeBool(true)
	b.Clear()
	if b != (Bool{}) {
		t.Error("value should be zero")
	}
}

func TestBool_IsSet(t *testing.T) {
	{
		b := MakeBool(true)
		if !b.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var b Bool
		if b.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var b Bool
		b.Set(true)
		if !b.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		b := MakeBool(true)
		b.Clear()
		if b.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestByte_Get(t *testing.T) {
	{
		b := MakeByte('b')
		v, ok := b.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 'b' {
			t.Error("value should be expected")
		}
	}
	{
		var b Byte
		v, ok := b.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != '\000' {
			t.Error("value should be zero")
		}
	}
	{
		var b Byte
		b.Set('b')
		v, ok := b.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 'b' {
			t.Error("value should be expected")
		}
	}
	{
		b := MakeByte('b')
		b.Clear()
		v, ok := b.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != '\000' {
			t.Error("value should be zero")
		}
	}
}

func TestByte_Set(t *testing.T) {
	var b Byte
	b.Set('b')
	if b != MakeByte('b') {
		t.Error("value should be expected")
	}
}

func TestByte_Clear(t *testing.T) {
	b := MakeByte('b')
	b.Clear()
	if b != (Byte{}) {
		t.Error("value should be zero")
	}
}

func TestByte_IsSet(t *testing.T) {
	{
		b := MakeByte('b')
		if !b.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var b Byte
		if b.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var b Byte
		b.Set('b')
		if !b.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		b := MakeByte('b')
		b.Clear()
		if b.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestInt_Get(t *testing.T) {
	{
		i := MakeInt(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var i Int
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var i Int
		i.Set(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		i := MakeInt(100)
		i.Clear()
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestInt_Set(t *testing.T) {
	var i Int
	i.Set(100)
	if i != MakeInt(100) {
		t.Error("value should be expected")
	}
}

func TestInt_Clear(t *testing.T) {
	i := MakeInt(100)
	i.Clear()
	if i != (Int{}) {
		t.Error("value should be zero")
	}
}

func TestInt_IsSet(t *testing.T) {
	{
		i := MakeInt(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var i Int
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var i Int
		i.Set(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		i := MakeInt(100)
		i.Clear()
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUint_Get(t *testing.T) {
	{
		u := MakeUint(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uint
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uint
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUint(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUint_Set(t *testing.T) {
	var u Uint
	u.Set(100)
	if u != MakeUint(100) {
		t.Error("value should be expected")
	}
}

func TestUint_Clear(t *testing.T) {
	u := MakeUint(100)
	u.Clear()
	if u != (Uint{}) {
		t.Error("value should be zero")
	}
}

func TestUint_IsSet(t *testing.T) {
	{
		u := MakeUint(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uint
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uint
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUint(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestInt8_Get(t *testing.T) {
	{
		i := MakeInt8(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var i Int8
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var i Int8
		i.Set(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		i := MakeInt8(100)
		i.Clear()
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestInt8_Set(t *testing.T) {
	var i Int8
	i.Set(100)
	if i != MakeInt8(100) {
		t.Error("value should be expected")
	}
}

func TestInt8_Clear(t *testing.T) {
	i := MakeInt8(100)
	i.Clear()
	if i != (Int8{}) {
		t.Error("value should be zero")
	}
}

func TestInt8_IsSet(t *testing.T) {
	{
		i := MakeInt8(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var i Int8
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var i Int8
		i.Set(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		i := MakeInt8(100)
		i.Clear()
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUint8_Get(t *testing.T) {
	{
		u := MakeUint8(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uint8
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uint8
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUint8(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUint8_Set(t *testing.T) {
	var u Uint8
	u.Set(100)
	if u != MakeUint8(100) {
		t.Error("value should be expected")
	}
}

func TestUint8_Clear(t *testing.T) {
	u := MakeUint8(100)
	u.Clear()
	if u != (Uint8{}) {
		t.Error("value should be zero")
	}
}

func TestUint8_IsSet(t *testing.T) {
	{
		u := MakeUint8(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uint8
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uint8
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUint8(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestInt16_Get(t *testing.T) {
	{
		i := MakeInt16(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var i Int16
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var i Int16
		i.Set(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		i := MakeInt16(100)
		i.Clear()
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestInt16_Set(t *testing.T) {
	var i Int16
	i.Set(100)
	if i != MakeInt16(100) {
		t.Error("value should be expected")
	}
}

func TestInt16_Clear(t *testing.T) {
	i := MakeInt16(100)
	i.Clear()
	if i != (Int16{}) {
		t.Error("value should be zero")
	}
}

func TestInt16_IsSet(t *testing.T) {
	{
		i := MakeInt16(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var i Int16
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var i Int16
		i.Set(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		i := MakeInt16(100)
		i.Clear()
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUint16_Get(t *testing.T) {
	{
		u := MakeUint16(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uint16
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uint16
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUint16(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUint16_Set(t *testing.T) {
	var u Uint16
	u.Set(100)
	if u != MakeUint16(100) {
		t.Error("value should be expected")
	}
}

func TestUint16_Clear(t *testing.T) {
	u := MakeUint16(100)
	u.Clear()
	if u != (Uint16{}) {
		t.Error("value should be zero")
	}
}

func TestUint16_IsSet(t *testing.T) {
	{
		u := MakeUint16(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uint16
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uint16
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUint16(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestInt32_Get(t *testing.T) {
	{
		i := MakeInt32(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var i Int32
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var i Int32
		i.Set(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		i := MakeInt32(100)
		i.Clear()
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestInt32_Set(t *testing.T) {
	var i Int32
	i.Set(100)
	if i != MakeInt32(100) {
		t.Error("value should be expected")
	}
}

func TestInt32_Clear(t *testing.T) {
	i := MakeInt32(100)
	i.Clear()
	if i != (Int32{}) {
		t.Error("value should be zero")
	}
}

func TestInt32_IsSet(t *testing.T) {
	{
		i := MakeInt32(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var i Int32
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var i Int32
		i.Set(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		i := MakeInt32(100)
		i.Clear()
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUint32_Get(t *testing.T) {
	{
		u := MakeUint32(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uint32
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uint32
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUint32(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUint32_Set(t *testing.T) {
	var u Uint32
	u.Set(100)
	if u != MakeUint32(100) {
		t.Error("value should be expected")
	}
}

func TestUint32_Clear(t *testing.T) {
	u := MakeUint32(100)
	u.Clear()
	if u != (Uint32{}) {
		t.Error("value should be zero")
	}
}

func TestUint32_IsSet(t *testing.T) {
	{
		u := MakeUint32(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uint32
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uint32
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUint32(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestInt64_Get(t *testing.T) {
	{
		i := MakeInt64(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var i Int64
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var i Int64
		i.Set(100)
		v, ok := i.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		i := MakeInt64(100)
		i.Clear()
		v, ok := i.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestInt64_Set(t *testing.T) {
	var i Int64
	i.Set(100)
	if i != MakeInt64(100) {
		t.Error("value should be expected")
	}
}

func TestInt64_Clear(t *testing.T) {
	i := MakeInt64(100)
	i.Clear()
	if i != (Int64{}) {
		t.Error("value should be zero")
	}
}

func TestInt64_IsSet(t *testing.T) {
	{
		i := MakeInt64(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var i Int64
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var i Int64
		i.Set(100)
		if !i.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		i := MakeInt64(100)
		i.Clear()
		if i.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUint64_Get(t *testing.T) {
	{
		u := MakeUint64(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uint64
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uint64
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUint64(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUint64_Set(t *testing.T) {
	var u Uint64
	u.Set(100)
	if u != MakeUint64(100) {
		t.Error("value should be expected")
	}
}

func TestUint64_Clear(t *testing.T) {
	u := MakeUint64(100)
	u.Clear()
	if u != (Uint64{}) {
		t.Error("value should be zero")
	}
}

func TestUint64_IsSet(t *testing.T) {
	{
		u := MakeUint64(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uint64
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uint64
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUint64(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestUintptr_Get(t *testing.T) {
	{
		u := MakeUintptr(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		var u Uintptr
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var u Uintptr
		u.Set(100)
		v, ok := u.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100 {
			t.Error("value should be expected")
		}
	}
	{
		u := MakeUintptr(100)
		u.Clear()
		v, ok := u.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestUintptr_Set(t *testing.T) {
	var u Uintptr
	u.Set(100)
	if u != MakeUintptr(100) {
		t.Error("value should be expected")
	}
}

func TestUintptr_Clear(t *testing.T) {
	u := MakeUintptr(100)
	u.Clear()
	if u != (Uintptr{}) {
		t.Error("value should be zero")
	}
}

func TestUintptr_IsSet(t *testing.T) {
	{
		u := MakeUintptr(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var u Uintptr
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var u Uintptr
		u.Set(100)
		if !u.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		u := MakeUintptr(100)
		u.Clear()
		if u.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestFloat32_Get(t *testing.T) {
	{
		f := MakeFloat32(1.0)
		v, ok := f.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 1.0 {
			t.Error("value should be expected")
		}
	}
	{
		var f Float32
		v, ok := f.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0.0 {
			t.Error("value should be zero")
		}
	}
	{
		var f Float32
		f.Set(1.0)
		v, ok := f.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 1.0 {
			t.Error("value should be expected")
		}
	}
	{
		f := MakeFloat32(1.0)
		f.Clear()
		v, ok := f.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0.0 {
			t.Error("value should be zero")
		}
	}
}

func TestFloat32_Set(t *testing.T) {
	var f Float32
	f.Set(1.0)
	if f != MakeFloat32(1.0) {
		t.Error("value should be expected")
	}
}

func TestFloat32_Clear(t *testing.T) {
	f := MakeFloat32(1.0)
	f.Clear()
	if f != (Float32{}) {
		t.Error("value should be zero")
	}
}

func TestFloat32_IsSet(t *testing.T) {
	{
		f := MakeFloat32(1.0)
		if !f.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var f Float32
		if f.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var f Float32
		f.Set(1.0)
		if !f.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		f := MakeFloat32(1.0)
		f.Clear()
		if f.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestFloat64_Get(t *testing.T) {
	{
		f := MakeFloat64(1.0)
		v, ok := f.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 1.0 {
			t.Error("value should be expected")
		}
	}
	{
		var f Float64
		v, ok := f.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0.0 {
			t.Error("value should be zero")
		}
	}
	{
		var f Float64
		f.Set(1.0)
		v, ok := f.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 1.0 {
			t.Error("value should be expected")
		}
	}
	{
		f := MakeFloat64(1.0)
		f.Clear()
		v, ok := f.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0.0 {
			t.Error("value should be zero")
		}
	}
}

func TestFloat64_Set(t *testing.T) {
	var f Float64
	f.Set(1.0)
	if f != MakeFloat64(1.0) {
		t.Error("value should be expected")
	}
}

func TestFloat64_Clear(t *testing.T) {
	f := MakeFloat64(1.0)
	f.Clear()
	if f != (Float64{}) {
		t.Error("value should be zero")
	}
}

func TestFloat64_IsSet(t *testing.T) {
	{
		f := MakeFloat64(1.0)
		if !f.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var f Float64
		if f.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var f Float64
		f.Set(1.0)
		if !f.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		f := MakeFloat64(1.0)
		f.Clear()
		if f.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestComplex64_Get(t *testing.T) {
	{
		c := MakeComplex64(complex(1, 2))
		v, ok := c.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != complex(1, 2) {
			t.Error("value should be expected")
		}
	}
	{
		var c Complex64
		v, ok := c.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != complex(0, 0) {
			t.Error("value should be zero")
		}
	}
	{
		var c Complex64
		c.Set(complex(1, 2))
		v, ok := c.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != complex(1, 2) {
			t.Error("value should be expected")
		}
	}
	{
		c := MakeComplex64(complex(1, 2))
		c.Clear()
		v, ok := c.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != complex(0, 0) {
			t.Error("value should be zero")
		}
	}
}

func TestComplex64_Set(t *testing.T) {
	var c Complex64
	c.Set(complex(1, 2))
	if c != MakeComplex64(complex(1, 2)) {
		t.Error("value should be expected")
	}
}

func TestComplex64_Clear(t *testing.T) {
	c := MakeComplex64(complex(1, 2))
	c.Clear()
	if c != (Complex64{}) {
		t.Error("value should be zero")
	}
}

func TestComplex64_IsSet(t *testing.T) {
	{
		c := MakeComplex64(complex(1, 2))
		if !c.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var c Complex64
		if c.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var c Complex64
		c.Set(complex(1, 2))
		if !c.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		c := MakeComplex64(complex(1, 2))
		c.Clear()
		if c.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestComplex128_Get(t *testing.T) {
	{
		c := MakeComplex128(complex(1, 2))
		v, ok := c.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != complex(1, 2) {
			t.Error("value should be expected")
		}
	}
	{
		var c Complex128
		v, ok := c.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != complex(0, 0) {
			t.Error("value should be zero")
		}
	}
	{
		var c Complex128
		c.Set(complex(1, 2))
		v, ok := c.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != complex(1, 2) {
			t.Error("value should be expected")
		}
	}
	{
		c := MakeComplex128(complex(1, 2))
		c.Clear()
		v, ok := c.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != complex(0, 0) {
			t.Error("value should be zero")
		}
	}
}

func TestComplex128_Set(t *testing.T) {
	var c Complex128
	c.Set(complex(1, 2))
	if c != MakeComplex128(complex(1, 2)) {
		t.Error("value should be expected")
	}
}

func TestComplex128_Clear(t *testing.T) {
	c := MakeComplex128(complex(1, 2))
	c.Clear()
	if c != (Complex128{}) {
		t.Error("value should be zero")
	}
}

func TestComplex128_IsSet(t *testing.T) {
	{
		c := MakeComplex128(complex(1, 2))
		if !c.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var c Complex128
		if c.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var c Complex128
		c.Set(complex(1, 2))
		if !c.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		c := MakeComplex128(complex(1, 2))
		c.Clear()
		if c.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestRune_Get(t *testing.T) {
	{
		r := MakeRune('😊')
		v, ok := r.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != '😊' {
			t.Error("value should be expected")
		}
	}
	{
		var r Rune
		v, ok := r.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != '\000' {
			t.Error("value should be zero")
		}
	}
	{
		var r Rune
		r.Set('😊')
		v, ok := r.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != '😊' {
			t.Error("value should be expected")
		}
	}
	{
		r := MakeRune('😊')
		r.Clear()
		v, ok := r.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != '\000' {
			t.Error("value should be zero")
		}
	}
}

func TestRune_Set(t *testing.T) {
	var r Rune
	r.Set('😊')
	if r != MakeRune('😊') {
		t.Error("value should be expected")
	}
}

func TestRune_Clear(t *testing.T) {
	r := MakeRune('😊')
	r.Clear()
	if r != (Rune{}) {
		t.Error("value should be zero")
	}
}

func TestRune_IsSet(t *testing.T) {
	{
		r := MakeRune('😊')
		if !r.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var r Rune
		if r.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var r Rune
		r.Set('😊')
		if !r.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		r := MakeRune('😊')
		r.Clear()
		if r.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestString_Get(t *testing.T) {
	{
		s := MakeString("foo")
		v, ok := s.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != "foo" {
			t.Error("value should be expected")
		}
	}
	{
		var s String
		v, ok := s.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != "" {
			t.Error("value should be zero")
		}
	}
	{
		var s String
		s.Set("foo")
		v, ok := s.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != "foo" {
			t.Error("value should be expected")
		}
	}
	{
		s := MakeString("foo")
		s.Clear()
		v, ok := s.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != "" {
			t.Error("value should be zero")
		}
	}
}

func TestString_Set(t *testing.T) {
	var s String
	s.Set("foo")
	if s != MakeString("foo") {
		t.Error("value should be expected")
	}
}

func TestString_Clear(t *testing.T) {
	s := MakeString("foo")
	s.Clear()
	if s != (String{}) {
		t.Error("value should be zero")
	}
}

func TestString_IsSet(t *testing.T) {
	{
		s := MakeString("foo")
		if !s.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var s String
		if s.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var s String
		s.Set("foo")
		if !s.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		s := MakeString("foo")
		s.Clear()
		if s.IsSet() {
			t.Error("value should be unset")
		}
	}
}

func TestDuration_Get(t *testing.T) {
	{
		d := MakeDuration(100 * time.Second)
		v, ok := d.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100*time.Second {
			t.Error("value should be expected")
		}
	}
	{
		var d Duration
		v, ok := d.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
	{
		var d Duration
		d.Set(100 * time.Second)
		v, ok := d.Get()
		if !ok {
			t.Error("value should be set")
		}
		if v != 100*time.Second {
			t.Error("value should be expected")
		}
	}
	{
		d := MakeDuration(100 * time.Second)
		d.Clear()
		v, ok := d.Get()
		if ok {
			t.Error("value should be unset")
		}
		if v != 0 {
			t.Error("value should be zero")
		}
	}
}

func TestDuration_Set(t *testing.T) {
	var d Duration
	d.Set(100 * time.Second)
	if d != MakeDuration(100*time.Second) {
		t.Error("value should be expected")
	}
}

func TestDuration_Clear(t *testing.T) {
	d := MakeDuration(100 * time.Second)
	d.Clear()
	if d != (Duration{}) {
		t.Error("value should be zero")
	}
}

func TestDuration_IsSet(t *testing.T) {
	{
		d := MakeDuration(100 * time.Second)
		if !d.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		var d Duration
		if d.IsSet() {
			t.Error("value should be unset")
		}
	}
	{
		var d Duration
		d.Set(100 * time.Second)
		if !d.IsSet() {
			t.Error("value should be set")
		}
	}
	{
		d := MakeDuration(100 * time.Second)
		d.Clear()
		if d.IsSet() {
			t.Error("value should be unset")
		}
	}
}
