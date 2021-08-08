// Code generated. DO NOT EDIT.

package optional

import (
	"encoding/json"
)

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(b.value)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*b = Bool{}
		return nil
	}
	return json.Unmarshal(data, &b.value)
}

func (b Byte) MarshalJSON() ([]byte, error) {
	if !b.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(b.value)
}

func (b *Byte) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*b = Byte{}
		return nil
	}
	return json.Unmarshal(data, &b.value)
}

func (i Int) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = Int{}
		return nil
	}
	return json.Unmarshal(data, &i.value)
}

func (u Uint) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uint{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (i Int8) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int8) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = Int8{}
		return nil
	}
	return json.Unmarshal(data, &i.value)
}

func (u Uint8) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint8) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uint8{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = Int16{}
		return nil
	}
	return json.Unmarshal(data, &i.value)
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint16) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uint16{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = Int32{}
		return nil
	}
	return json.Unmarshal(data, &i.value)
}

func (u Uint32) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uint32{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(i.value)
}

func (i *Int64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*i = Int64{}
		return nil
	}
	return json.Unmarshal(data, &i.value)
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uint64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uint64{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (u Uintptr) MarshalJSON() ([]byte, error) {
	if !u.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(u.value)
}

func (u *Uintptr) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Uintptr{}
		return nil
	}
	return json.Unmarshal(data, &u.value)
}

func (f Float32) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

func (f *Float32) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*f = Float32{}
		return nil
	}
	return json.Unmarshal(data, &f.value)
}

func (f Float64) MarshalJSON() ([]byte, error) {
	if !f.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(f.value)
}

func (f *Float64) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*f = Float64{}
		return nil
	}
	return json.Unmarshal(data, &f.value)
}

func (r Rune) MarshalJSON() ([]byte, error) {
	if !r.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(r.value)
}

func (r *Rune) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*r = Rune{}
		return nil
	}
	return json.Unmarshal(data, &r.value)
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(s.value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = String{}
		return nil
	}
	return json.Unmarshal(data, &s.value)
}

func (d Duration) MarshalJSON() ([]byte, error) {
	if !d.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal(d.value)
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*d = Duration{}
		return nil
	}
	return json.Unmarshal(data, &d.value)
}
