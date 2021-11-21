//go:generate go run $GOFILE

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

const packagePath = "github.com/go-tk/optional"
const packageName = "optional"

var typeInfos = []struct {
	FullName        string
	CapitalizedName string
	ShortenName     string
	ZeroLiteral     string
	NonZeroLiteral  string
}{
	{"bool", "Bool", "b", "false", "true"},
	{"byte", "Byte", "b", `'\000'`, "'b'"},
	{"int", "Int", "i", "0", "100"},
	{"uint", "Uint", "u", "0", "100"},
	{"int8", "Int8", "i", "0", "100"},
	{"uint8", "Uint8", "u", "0", "100"},
	{"int16", "Int16", "i", "0", "100"},
	{"uint16", "Uint16", "u", "0", "100"},
	{"int32", "Int32", "i", "0", "100"},
	{"uint32", "Uint32", "u", "0", "100"},
	{"int64", "Int64", "i", "0", "100"},
	{"uint64", "Uint64", "u", "0", "100"},
	{"uintptr", "Uintptr", "u", "0", "100"},
	{"float32", "Float32", "f", "0.0", "1.0"},
	{"float64", "Float64", "f", "0.0", "1.0"},
	{"complex64", "Complex64", "c", "complex(0, 0)", "complex(1, 2)"},
	{"complex128", "Complex128", "c", "complex(0, 0)", "complex(1, 2)"},
	{"rune", "Rune", "r", `'\000'`, "'😊'"},
	{"string", "String", "s", `""`, `"foo"`},
	{"time.Duration", "Duration", "d", "0", "100 * time.Second"},
}

func main() {
	fileName2Code := map[string][]byte{
		packageName + ".go":      generateCode(),
		packageName + "_test.go": generateTestingCode(),
		"json.go":                generateCodeForJSON(),
		"json_test.go":           generateTestingCodeForJSON(),
		"yaml.go":                generateCodeForYAML(),
		"yaml_test.go":           generateTestingCodeForYAML(),
	}
	for fileName, code := range fileName2Code {
		code2, err := format.Source(code)
		if err == nil {
			code = code2
		}
		filePath := filepath.Join("..", fileName)
		if err2 := ioutil.WriteFile(filePath, code, 0644); err2 != nil {
			panic(err2)
		}
		if err != nil {
			panic(err)
		}
	}
}

func generateCode() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`

import (
	"time"
)
{{- range .}}

type {{.CapitalizedName}} struct {
       value {{.FullName}}
       hasValue bool
}

func Make{{.CapitalizedName}}(value {{.FullName}}) {{.CapitalizedName}} {
       return {{.CapitalizedName}}{value, true}
}

func ({{.ShortenName}} *{{.CapitalizedName}}) Set(value {{.FullName}}) {
	*{{.ShortenName}} = Make{{.CapitalizedName}}(value)
}

func ({{.ShortenName}} *{{.CapitalizedName}}) Clear() {
	*{{.ShortenName}} = {{.CapitalizedName}}{}
}

func ({{.ShortenName}} {{.CapitalizedName}}) Value() {{.FullName}} { return {{.ShortenName}}.value }
func ({{.ShortenName}} {{.CapitalizedName}}) HasValue() bool { return {{.ShortenName}}.hasValue }
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func generateTestingCode() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`_test

import (
	"testing"
	"time"

	. "`+packagePath+`"
)
{{- range .}}

func Test{{.CapitalizedName}}_Set(t *testing.T) {
	var {{.ShortenName}} {{.CapitalizedName}}
	{{.ShortenName}}.Set({{.NonZeroLiteral}})
	if {{.ShortenName}} != Make{{.CapitalizedName}}({{.NonZeroLiteral}}) {
		t.Error("should be expected value")
	}
}

func Test{{.CapitalizedName}}_Clear(t *testing.T) {
	{{.ShortenName}} := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
	{{.ShortenName}}.Clear()
	if {{.ShortenName}} != ({{.CapitalizedName}}{}) {
		t.Error("should be zero value")
	}
}

func Test{{.CapitalizedName}}_Value(t *testing.T) {
	{
		{{.ShortenName}} := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		if {{.ShortenName}}.Value() != {{.NonZeroLiteral}} {
			t.Error("should be expected value")
		}
	}
	{
		var {{.ShortenName}} {{.CapitalizedName}}
		if {{.ShortenName}}.Value() != {{.ZeroLiteral}} {
			t.Error("should be zero value")
		}
	}
	{
		var {{.ShortenName}} {{.CapitalizedName}}
		{{.ShortenName}}.Set({{.NonZeroLiteral}})
		if {{.ShortenName}}.Value() != {{.NonZeroLiteral}} {
			t.Error("should be expected value")
		}
	}
	{
		{{.ShortenName}} := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		{{.ShortenName}}.Clear()
		if {{.ShortenName}}.Value() != {{.ZeroLiteral}} {
			t.Error("should be zero value")
		}
	}
}

func Test{{.CapitalizedName}}_HasValue(t *testing.T) {
	{
		{{.ShortenName}} := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		if !{{.ShortenName}}.HasValue() {
			t.Error("should have value")
		}
	}
	{
		var {{.ShortenName}} {{.CapitalizedName}}
		if {{.ShortenName}}.HasValue() {
			t.Error("should have no value")
		}
	}
	{
		var {{.ShortenName}} {{.CapitalizedName}}
		{{.ShortenName}}.Set({{.NonZeroLiteral}})
		if !{{.ShortenName}}.HasValue() {
			t.Error("should have value")
		}
	}
	{
		{{.ShortenName}} := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		{{.ShortenName}}.Clear()
		if {{.ShortenName}}.HasValue() {
			t.Error("should have no value")
		}
	}
}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func generateCodeForJSON() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`

import (
	"encoding/json"
	"fmt"
	"time"
)
{{- range .}}
{{- if and (ne .FullName "complex64") (ne .FullName "complex128")}}
{{- if eq .FullName "time.Duration"}}

func ({{.ShortenName}} {{.CapitalizedName}}) MarshalJSON() ([]byte, error) {
	if !{{.ShortenName}}.hasValue {
		return []byte("null"), nil
	}
	valueStr := {{.ShortenName}}.value.String()
	return json.Marshal(valueStr)
}

func ({{.ShortenName}} *{{.CapitalizedName}}) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		{{.ShortenName}}.Clear()
		return nil
	}
	var valueStr string
	if err := json.Unmarshal(data, &valueStr); err != nil {
		return err
	}
	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return fmt.Errorf("duration parse failed; valueStr=%q: %v", valueStr, err)
	}
	{{.ShortenName}}.Set(value)
	return nil
}
{{- else}}

func ({{.ShortenName}} {{.CapitalizedName}}) MarshalJSON() ([]byte, error) {
	if !{{.ShortenName}}.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal({{.ShortenName}}.value)
}

func ({{.ShortenName}} *{{.CapitalizedName}}) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		{{.ShortenName}}.Clear()
		return nil
	}
	var value {{.FullName}}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	{{.ShortenName}}.Set(value)
	return nil
}
{{- end}}
{{- end}}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func generateTestingCodeForJSON() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`_test

import (
	"encoding/json"
	"testing"
	"time"

	. "`+packagePath+`"
)
{{- range .}}
{{- if and (ne .FullName "complex64") (ne .FullName "complex128")}}

func Test{{.CapitalizedName}}_MarshalAndUnmarshalJSON(t *testing.T) {
	{
		var {{.ShortenName}}1 {{.CapitalizedName}}
		data, err := json.Marshal({{.ShortenName}}1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null" {
			t.Error("should be null json")
		}
		{{.ShortenName}}2 := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		err = json.Unmarshal(data, &{{.ShortenName}}2)
		if err != nil {
			t.Fatal(err)
		}
		if {{.ShortenName}}1 != {{.ShortenName}}2 {
			t.Error("should be equal values")
		}
	}
	{
		{{.ShortenName}}1 := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		data, err := json.Marshal({{.ShortenName}}1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null" {
			t.Error("should not be null json")
		}
		var {{.ShortenName}}2 {{.CapitalizedName}}
		err = json.Unmarshal(data, &{{.ShortenName}}2)
		if err != nil {
			t.Fatal(err)
		}
		if {{.ShortenName}}1 != {{.ShortenName}}2 {
			t.Error("should be equal values")
		}
	}
}

func Test{{.CapitalizedName}}_UnmarshalJSON(t *testing.T) {
	var {{.ShortenName}} {{.CapitalizedName}}
	err := json.Unmarshal([]byte("{}"), &{{.ShortenName}})
	if err == nil {
		t.Error("should encounter error")
	}
}
{{- end}}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func generateCodeForYAML() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`

import (
	"fmt"
	"time"
)
{{- range .}}
{{- if and (ne .FullName "complex64") (ne .FullName "complex128")}}
{{- if eq .FullName "time.Duration"}}

func ({{.ShortenName}} {{.CapitalizedName}}) MarshalYAML() (interface{}, error) {
	if !{{.ShortenName}}.hasValue {
		return nil, nil
	}
	valueStr := {{.ShortenName}}.value.String()
	return valueStr, nil
}

func ({{.ShortenName}} *{{.CapitalizedName}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var valueStr *string
	if err := unmarshal(&valueStr); err != nil {
		return err
	}
	if valueStr == nil {
		{{.ShortenName}}.Clear()
		return nil
	}
	value, err := time.ParseDuration(*valueStr)
	if err != nil {
		return fmt.Errorf("duration parse failed; valueStr=%q: %v", *valueStr, err)
	}
	{{.ShortenName}}.Set(value)
	return nil
}
{{- else}}

func ({{.ShortenName}} {{.CapitalizedName}}) MarshalYAML() (interface{}, error) {
	if !{{.ShortenName}}.hasValue {
		return nil, nil
	}
	return {{.ShortenName}}.value, nil
}

func ({{.ShortenName}} *{{.CapitalizedName}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value *{{.FullName}}
	if err := unmarshal(&value); err != nil {
		return err
	}
	if value == nil {
		{{.ShortenName}}.Clear()
		return nil
	}
	{{.ShortenName}}.Set(*value)
	return nil
}
{{- end}}
{{- end}}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func generateTestingCodeForYAML() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`_test

import (
	"testing"
	"time"

	. "`+packagePath+`"
	"gopkg.in/yaml.v2"
)
{{- range .}}
{{- if and (ne .FullName "complex64") (ne .FullName "complex128")}}

func Test{{.CapitalizedName}}_MarshalAndUnmarshalYAML(t *testing.T) {
	{
		var {{.ShortenName}}1 {{.CapitalizedName}}
		data, err := yaml.Marshal({{.ShortenName}}1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "null\n" {
			t.Error("should be null yaml")
		}
		{{.ShortenName}}2 := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		err = yaml.Unmarshal(data, &{{.ShortenName}}2)
		if err != nil {
			t.Fatal(err)
		}
		if {{.ShortenName}}1 != {{.ShortenName}}2 {
			t.Error("should be equal values")
		}
	}
	{
		{{.ShortenName}}1 := Make{{.CapitalizedName}}({{.NonZeroLiteral}})
		data, err := yaml.Marshal({{.ShortenName}}1)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) == "null\n" {
			t.Error("should not be null yaml")
		}
		var {{.ShortenName}}2 {{.CapitalizedName}}
		err = yaml.Unmarshal(data, &{{.ShortenName}}2)
		if err != nil {
			t.Fatal(err)
		}
		if {{.ShortenName}}1 != {{.ShortenName}}2 {
			t.Error("should be equal values")
		}
	}
}

func Test{{.CapitalizedName}}_UnmarshalYAML(t *testing.T) {
	var {{.ShortenName}} {{.CapitalizedName}}
	err := yaml.Unmarshal([]byte("{}"), &{{.ShortenName}})
	if err == nil {
		t.Error("should encounter error")
	}
}
{{- end}}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
