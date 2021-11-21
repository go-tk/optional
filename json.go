// Code generated. DO NOT EDIT.

package optional

import (
	"encoding/json"
	"fmt"
	"time"
)

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(b.value)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Clear()
		return nil
	}
	var value bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	b.Set(value)
	return nil
}

func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(b.value)
}

func (b *Byte) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		b.Clear()
		return nil
	}
	var value byte
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	b.Set(value)
	return nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Clear()
		return nil
	}
	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.Set(value)
	return nil
}

func (u Uint) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uint
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (i Int8) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int8) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Clear()
		return nil
	}
	var value int8
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.Set(value)
	return nil
}

func (u Uint8) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint8) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uint8
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Clear()
		return nil
	}
	var value int16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.Set(value)
	return nil
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uint16
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Clear()
		return nil
	}
	var value int32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.Set(value)
	return nil
}

func (u Uint32) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uint32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		i.Clear()
		return nil
	}
	var value int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	i.Set(value)
	return nil
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uint64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (u Uintptr) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uintptr) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		u.Clear()
		return nil
	}
	var value uintptr
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	u.Set(value)
	return nil
}

func (f Float32) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

func (f *Float32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Clear()
		return nil
	}
	var value float32
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	f.Set(value)
	return nil
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		f.Clear()
		return nil
	}
	var value float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	f.Set(value)
	return nil
}

func (r Rune) MarshalJSON() ([]byte, error) {
	if !r.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(r.value)
}

func (r *Rune) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		r.Clear()
		return nil
	}
	var value rune
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	r.Set(value)
	return nil
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(s.value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.Clear()
		return nil
	}
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	s.Set(value)
	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	if !d.hasValue {
		return []byte("null"), nil
	}
	valueStr := d.value.String()
	return json.Marshal(valueStr)
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		d.Clear()
		return nil
	}
	var valueStr string
	if err := json.Unmarshal(data, &valueStr); err != nil {
		return err
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return fmt.Errorf("optional: duration parse failed; valueStr=%q: %v", valueStr, err)
	}
	d.Set(value)
	return nil
}
