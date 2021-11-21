// Code generated. DO NOT EDIT.

package optional

import (
	"fmt"
	"time"
)

func (b Bool) MarshalYAML() (interface{}, error) {
	if !b.hasValue {
		return nil, nil
	}
	return b.value, nil
}

func (b *Bool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *bool
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		b.Clear()
		return nil
	}
	b.Set(*value)
	return nil
}

func (b Byte) MarshalYAML() (interface{}, error) {
	if !b.hasValue {
		return nil, nil
	}
	return b.value, nil
}

func (b *Byte) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *byte
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		b.Clear()
		return nil
	}
	b.Set(*value)
	return nil
}

func (i Int) MarshalYAML() (interface{}, error) {
	if !i.hasValue {
		return nil, nil
	}
	return i.value, nil
}

func (i *Int) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *int
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		i.Clear()
		return nil
	}
	i.Set(*value)
	return nil
}

func (u Uint) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uint
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (i Int8) MarshalYAML() (interface{}, error) {
	if !i.hasValue {
		return nil, nil
	}
	return i.value, nil
}

func (i *Int8) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *int8
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		i.Clear()
		return nil
	}
	i.Set(*value)
	return nil
}

func (u Uint8) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uint8) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uint8
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (i Int16) MarshalYAML() (interface{}, error) {
	if !i.hasValue {
		return nil, nil
	}
	return i.value, nil
}

func (i *Int16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *int16
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		i.Clear()
		return nil
	}
	i.Set(*value)
	return nil
}

func (u Uint16) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uint16) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uint16
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (i Int32) MarshalYAML() (interface{}, error) {
	if !i.hasValue {
		return nil, nil
	}
	return i.value, nil
}

func (i *Int32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *int32
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		i.Clear()
		return nil
	}
	i.Set(*value)
	return nil
}

func (u Uint32) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uint32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uint32
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (i Int64) MarshalYAML() (interface{}, error) {
	if !i.hasValue {
		return nil, nil
	}
	return i.value, nil
}

func (i *Int64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *int64
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		i.Clear()
		return nil
	}
	i.Set(*value)
	return nil
}

func (u Uint64) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uint64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uint64
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (u Uintptr) MarshalYAML() (interface{}, error) {
	if !u.hasValue {
		return nil, nil
	}
	return u.value, nil
}

func (u *Uintptr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *uintptr
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		u.Clear()
		return nil
	}
	u.Set(*value)
	return nil
}

func (f Float32) MarshalYAML() (interface{}, error) {
	if !f.hasValue {
		return nil, nil
	}
	return f.value, nil
}

func (f *Float32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *float32
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		f.Clear()
		return nil
	}
	f.Set(*value)
	return nil
}

func (f Float64) MarshalYAML() (interface{}, error) {
	if !f.hasValue {
		return nil, nil
	}
	return f.value, nil
}

func (f *Float64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *float64
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		f.Clear()
		return nil
	}
	f.Set(*value)
	return nil
}

func (r Rune) MarshalYAML() (interface{}, error) {
	if !r.hasValue {
		return nil, nil
	}
	return r.value, nil
}

func (r *Rune) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *rune
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		r.Clear()
		return nil
	}
	r.Set(*value)
	return nil
}

func (s String) MarshalYAML() (interface{}, error) {
	if !s.hasValue {
		return nil, nil
	}
	return s.value, nil
}

func (s *String) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *string
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		s.Clear()
		return nil
	}
	s.Set(*value)
	return nil
}

func (d Duration) MarshalYAML() (interface{}, error) {
	if !d.hasValue {
		return nil, nil
	}
	valueStr := d.value.String()
	return valueStr, nil
}

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var valueStr *string
	if err := unmarshal(&valueStr); err != nil {
		return err
	}
	if valueStr == nil {
		d.Clear()
		return nil
	}
	value, err := time.ParseDuration(*valueStr)
	if err != nil {
		return fmt.Errorf("duration parse failed; valueStr=%q: %v", *valueStr, err)
	}
	d.Set(value)
	return nil
}
