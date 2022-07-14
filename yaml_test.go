// Code generated. DO NOT EDIT.

package optional_test

import (
	"testing"
	"time"

	. "github.com/go-tk/optional"
	"gopkg.in/yaml.v2"
)

func TestBool_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var b1 Bool
		data, err := yaml.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		b2 := MakeBool(true)
		err = yaml.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
	{
		b1 := MakeBool(true)
		data, err := yaml.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var b2 Bool
		err = yaml.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
}

func TestBool_UnmarshalYAML(t *testing.T) {
	var b Bool
	err := yaml.Unmarshal([]byte("{}"), &b)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestByte_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var b1 Byte
		data, err := yaml.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		b2 := MakeByte('b')
		err = yaml.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
	{
		b1 := MakeByte('b')
		data, err := yaml.Marshal(b1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var b2 Byte
		err = yaml.Unmarshal(data, &b2)
		if err != nil {
			t.Fatal(err)
		}
		if b1 != b2 {
			t.Error("values should be equal")
		}
	}
}

func TestByte_UnmarshalYAML(t *testing.T) {
	var b Byte
	err := yaml.Unmarshal([]byte("{}"), &b)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var i1 Int
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		i2 := MakeInt(100)
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt(100)
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var i2 Int
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt_UnmarshalYAML(t *testing.T) {
	var i Int
	err := yaml.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uint
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUint(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uint
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint_UnmarshalYAML(t *testing.T) {
	var u Uint
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt8_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var i1 Int8
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		i2 := MakeInt8(100)
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt8(100)
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var i2 Int8
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt8_UnmarshalYAML(t *testing.T) {
	var i Int8
	err := yaml.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint8_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uint8
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUint8(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint8(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uint8
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint8_UnmarshalYAML(t *testing.T) {
	var u Uint8
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt16_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var i1 Int16
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		i2 := MakeInt16(100)
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt16(100)
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var i2 Int16
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt16_UnmarshalYAML(t *testing.T) {
	var i Int16
	err := yaml.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint16_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uint16
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUint16(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint16(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uint16
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint16_UnmarshalYAML(t *testing.T) {
	var u Uint16
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt32_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var i1 Int32
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		i2 := MakeInt32(100)
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt32(100)
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var i2 Int32
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt32_UnmarshalYAML(t *testing.T) {
	var i Int32
	err := yaml.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint32_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uint32
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUint32(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint32(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uint32
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint32_UnmarshalYAML(t *testing.T) {
	var u Uint32
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestInt64_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var i1 Int64
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		i2 := MakeInt64(100)
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
	{
		i1 := MakeInt64(100)
		data, err := yaml.Marshal(i1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var i2 Int64
		err = yaml.Unmarshal(data, &i2)
		if err != nil {
			t.Fatal(err)
		}
		if i1 != i2 {
			t.Error("values should be equal")
		}
	}
}

func TestInt64_UnmarshalYAML(t *testing.T) {
	var i Int64
	err := yaml.Unmarshal([]byte("{}"), &i)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUint64_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uint64
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUint64(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUint64(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uint64
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUint64_UnmarshalYAML(t *testing.T) {
	var u Uint64
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestUintptr_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var u1 Uintptr
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		u2 := MakeUintptr(100)
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
	{
		u1 := MakeUintptr(100)
		data, err := yaml.Marshal(u1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var u2 Uintptr
		err = yaml.Unmarshal(data, &u2)
		if err != nil {
			t.Fatal(err)
		}
		if u1 != u2 {
			t.Error("values should be equal")
		}
	}
}

func TestUintptr_UnmarshalYAML(t *testing.T) {
	var u Uintptr
	err := yaml.Unmarshal([]byte("{}"), &u)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestFloat32_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var f1 Float32
		data, err := yaml.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		f2 := MakeFloat32(1.0)
		err = yaml.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
	{
		f1 := MakeFloat32(1.0)
		data, err := yaml.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var f2 Float32
		err = yaml.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
}

func TestFloat32_UnmarshalYAML(t *testing.T) {
	var f Float32
	err := yaml.Unmarshal([]byte("{}"), &f)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestFloat64_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var f1 Float64
		data, err := yaml.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		f2 := MakeFloat64(1.0)
		err = yaml.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
	{
		f1 := MakeFloat64(1.0)
		data, err := yaml.Marshal(f1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var f2 Float64
		err = yaml.Unmarshal(data, &f2)
		if err != nil {
			t.Fatal(err)
		}
		if f1 != f2 {
			t.Error("values should be equal")
		}
	}
}

func TestFloat64_UnmarshalYAML(t *testing.T) {
	var f Float64
	err := yaml.Unmarshal([]byte("{}"), &f)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestRune_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var r1 Rune
		data, err := yaml.Marshal(r1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		r2 := MakeRune('😊')
		err = yaml.Unmarshal(data, &r2)
		if err != nil {
			t.Fatal(err)
		}
		if r1 != r2 {
			t.Error("values should be equal")
		}
	}
	{
		r1 := MakeRune('😊')
		data, err := yaml.Marshal(r1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var r2 Rune
		err = yaml.Unmarshal(data, &r2)
		if err != nil {
			t.Fatal(err)
		}
		if r1 != r2 {
			t.Error("values should be equal")
		}
	}
}

func TestRune_UnmarshalYAML(t *testing.T) {
	var r Rune
	err := yaml.Unmarshal([]byte("{}"), &r)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestString_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var s1 String
		data, err := yaml.Marshal(s1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		s2 := MakeString("foo")
		err = yaml.Unmarshal(data, &s2)
		if err != nil {
			t.Fatal(err)
		}
		if s1 != s2 {
			t.Error("values should be equal")
		}
	}
	{
		s1 := MakeString("foo")
		data, err := yaml.Marshal(s1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var s2 String
		err = yaml.Unmarshal(data, &s2)
		if err != nil {
			t.Fatal(err)
		}
		if s1 != s2 {
			t.Error("values should be equal")
		}
	}
}

func TestString_UnmarshalYAML(t *testing.T) {
	var s String
	err := yaml.Unmarshal([]byte("{}"), &s)
	if err == nil {
		t.Error("should encounter error")
	}
}

func TestDuration_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var d1 Duration
		data, err := yaml.Marshal(d1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("yaml should be null")
		}
		d2 := MakeDuration(100 * time.Second)
		err = yaml.Unmarshal(data, &d2)
		if err != nil {
			t.Fatal(err)
		}
		if d1 != d2 {
			t.Error("values should be equal")
		}
	}
	{
		d1 := MakeDuration(100 * time.Second)
		data, err := yaml.Marshal(d1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("yaml should not be null")
		}
		var d2 Duration
		err = yaml.Unmarshal(data, &d2)
		if err != nil {
			t.Fatal(err)
		}
		if d1 != d2 {
			t.Error("values should be equal")
		}
	}
}

func TestDuration_UnmarshalYAML(t *testing.T) {
	var d Duration
	err := yaml.Unmarshal([]byte("{}"), &d)
	if err == nil {
		t.Error("should encounter error")
	}
}
