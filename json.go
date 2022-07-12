// Code generated. DO NOT EDIT.

package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// MarshalJSON implements json.Marshaler.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(b.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Bool) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(b.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *Byte) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uint) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uint) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (i Int8) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int8) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uint8) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uint8) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int16) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uint16) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uint16) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int32) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uint32) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uint32) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(i.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uint64) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uint64) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (u Uintptr) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(u.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *Uintptr) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (f Float32) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Float32) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(f.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (f *Float64) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (r Rune) MarshalJSON() ([]byte, error) {
	if !r.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(r.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (r *Rune) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.hasValue {
		return bytesOfNull, nil
	}
	return json.Marshal(s.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *String) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
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

// MarshalJSON implements json.Marshaler.
func (d Duration) MarshalJSON() ([]byte, error) {
	if !d.hasValue {
		return bytesOfNull, nil
	}
	valueStr := d.value.String()
	return json.Marshal(valueStr)
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *Duration) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, bytesOfNull) == 0 {
		d.Clear()
		return nil
	}
	var valueStr string
	if err := json.Unmarshal(data, &valueStr); err != nil {
		return err
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return fmt.Errorf("parse duration; valueStr=%q: %w", valueStr, err)
	}
	d.Set(value)
	return nil
}

var bytesOfNull = []byte("null")
