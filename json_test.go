// Code generated. DO NOT EDIT.

package optional_test

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/go-tk/optional"
)

func TestBool_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var b1 Bool
		data, err := json.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		b2 := MakeBool(true)
		err = json.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
	{
		b1 := MakeBool(true)
		data, err := json.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var b2 Bool
		err = json.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
}

func TestBool_UnmarshalJSON(t *testing.T) {
	var b Bool
	err := json.Unmarshal([]byte("{}"), &b)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestByte_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var b1 Byte
		data, err := json.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		b2 := MakeByte('b')
		err = json.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
	{
		b1 := MakeByte('b')
		data, err := json.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var b2 Byte
		err = json.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
}

func TestByte_UnmarshalJSON(t *testing.T) {
	var b Byte
	err := json.Unmarshal([]byte("{}"), &b)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var i1 Int
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		i2 := MakeInt(100)
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt(100)
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var i2 Int
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	var i Int
	err := json.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uint
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUint(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uint
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint_UnmarshalJSON(t *testing.T) {
	var u Uint
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt8_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var i1 Int8
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		i2 := MakeInt8(100)
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt8(100)
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var i2 Int8
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt8_UnmarshalJSON(t *testing.T) {
	var i Int8
	err := json.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint8_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uint8
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUint8(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint8(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uint8
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint8_UnmarshalJSON(t *testing.T) {
	var u Uint8
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt16_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var i1 Int16
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		i2 := MakeInt16(100)
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt16(100)
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var i2 Int16
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt16_UnmarshalJSON(t *testing.T) {
	var i Int16
	err := json.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint16_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uint16
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUint16(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint16(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uint16
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint16_UnmarshalJSON(t *testing.T) {
	var u Uint16
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt32_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var i1 Int32
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		i2 := MakeInt32(100)
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt32(100)
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var i2 Int32
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt32_UnmarshalJSON(t *testing.T) {
	var i Int32
	err := json.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint32_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uint32
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUint32(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint32(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uint32
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint32_UnmarshalJSON(t *testing.T) {
	var u Uint32
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt64_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var i1 Int64
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		i2 := MakeInt64(100)
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt64(100)
		data, err := json.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var i2 Int64
		err = json.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt64_UnmarshalJSON(t *testing.T) {
	var i Int64
	err := json.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint64_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uint64
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUint64(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint64(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uint64
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint64_UnmarshalJSON(t *testing.T) {
	var u Uint64
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUintptr_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var u1 Uintptr
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		u2 := MakeUintptr(100)
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUintptr(100)
		data, err := json.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var u2 Uintptr
		err = json.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUintptr_UnmarshalJSON(t *testing.T) {
	var u Uintptr
	err := json.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestFloat32_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var f1 Float32
		data, err := json.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		f2 := MakeFloat32(1.0)
		err = json.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
	{
		f1 := MakeFloat32(1.0)
		data, err := json.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var f2 Float32
		err = json.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
}

func TestFloat32_UnmarshalJSON(t *testing.T) {
	var f Float32
	err := json.Unmarshal([]byte("{}"), &f)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestFloat64_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var f1 Float64
		data, err := json.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		f2 := MakeFloat64(1.0)
		err = json.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
	{
		f1 := MakeFloat64(1.0)
		data, err := json.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var f2 Float64
		err = json.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	var f Float64
	err := json.Unmarshal([]byte("{}"), &f)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestRune_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var r1 Rune
		data, err := json.Marshal(r1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		r2 := MakeRune('😊')
		err = json.Unmarshal(data, &r2)
		if err != nil {
			t.Fatal(err)
		}
		if r1 != r2 {
			t.Error("values should be equal")
		}
	}
	{
		r1 := MakeRune('😊')
		data, err := json.Marshal(r1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var r2 Rune
		err = json.Unmarshal(data, &r2)
		if err != nil {
			t.Fatal(err)
		}
		if r1 != r2 {
			t.Error("values should be equal")
		}
	}
}

func TestRune_UnmarshalJSON(t *testing.T) {
	var r Rune
	err := json.Unmarshal([]byte("{}"), &r)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestString_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var s1 String
		data, err := json.Marshal(s1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		s2 := MakeString("foo")
		err = json.Unmarshal(data, &s2)
		if err != nil {
			t.Fatal(err)
		}
		if s1 != s2 {
			t.Error("values should be equal")
		}
	}
	{
		s1 := MakeString("foo")
		data, err := json.Marshal(s1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var s2 String
		err = json.Unmarshal(data, &s2)
		if err != nil {
			t.Fatal(err)
		}
		if s1 != s2 {
			t.Error("values should be equal")
		}
	}
}

func TestString_UnmarshalJSON(t *testing.T) {
	var s String
	err := json.Unmarshal([]byte("{}"), &s)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestDuration_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var d1 Duration
		data, err := json.Marshal(d1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("json should be null")
		}
		d2 := MakeDuration(100 * time.Second)
		err = json.Unmarshal(data, &d2)
		if err != nil {
			t.Fatal(err)
		}
		if d1 != d2 {
			t.Error("values should be equal")
		}
	}
	{
		d1 := MakeDuration(100 * time.Second)
		data, err := json.Marshal(d1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("json should not be null")
		}
		var d2 Duration
		err = json.Unmarshal(data, &d2)
		if err != nil {
			t.Fatal(err)
		}
		if d1 != d2 {
			t.Error("values should be equal")
		}
	}
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	var d Duration
	err := json.Unmarshal([]byte("{}"), &d)
	if err == nil {
		t.Error("should encounter error")
	}
}
