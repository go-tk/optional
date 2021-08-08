//go:generate go run $GOFILE

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

const packageName = "optional"

var typeInfos = []struct {
	FullName        string
	CapitalizedName string
	ShortenName     string
}{
	{"bool", "Bool", "b"},
	{"byte", "Byte", "b"},
	{"int", "Int", "i"},
	{"uint", "Uint", "u"},
	{"int8", "Int8", "i"},
	{"uint8", "Uint8", "u"},
	{"int16", "Int16", "i"},
	{"uint16", "Uint16", "u"},
	{"int32", "Int32", "i"},
	{"uint32", "Uint32", "u"},
	{"int64", "Int64", "i"},
	{"uint64", "Uint64", "u"},
	{"uintptr", "Uintptr", "u"},
	{"float32", "Float32", "f"},
	{"float64", "Float64", "f"},
	{"complex64", "Complex64", "c"},
	{"complex128", "Complex128", "c"},
	{"rune", "Rune", "r"},
	{"string", "String", "s"},
	{"time.Duration", "Duration", "d"},
}

func main() {
	filePath2Code := map[string][]byte{
		filepath.Join("..", packageName+".go"): generateCode(),
		filepath.Join("..", "json.go"):         generateCodeForJSON(),
	}
	for filePath, code := range filePath2Code {
		code2, err := format.Source(code)
		if err == nil {
			code = code2
		}
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

func generateCodeForJSON() []byte {
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`

import (
	"encoding/json"
)
{{- range .}}
{{- if and (ne .FullName "complex64") (ne .FullName "complex128")}}

func ({{.ShortenName}} {{.CapitalizedName}}) MarshalJSON() ([]byte, error) {
	if !{{.ShortenName}}.hasValue {
		return []byte("null"), nil
	}
	return json.Marshal({{.ShortenName}}.value)
}

func ({{.ShortenName}} *{{.CapitalizedName}}) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*{{.ShortenName}} = {{.CapitalizedName}}{}
		return nil
	}
	return json.Unmarshal(data, &{{.ShortenName}}.value)
}
{{- end}}
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
