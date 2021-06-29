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

func main() {
	code1 := generateCode()
	code2, err := format.Source(code1)
	if err != nil {
		code2 = code1
	}
	fileName := filepath.Join("..", packageName+".go")
	if err := ioutil.WriteFile(fileName, code2, 0644); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
}

func generateCode() []byte {
	typeInfos := []struct {
		PackagePath     string
		ReferenceName   string
		CapitalizedName string
		ShortenName     string
	}{
		{"", "bool", "Bool", "b"},
		{"", "byte", "Byte", "b"},
		{"", "int", "Int", "i"},
		{"", "uint", "Uint", "u"},
		{"", "int8", "Int8", "i"},
		{"", "uint8", "Uint8", "u"},
		{"", "int16", "Int16", "i"},
		{"", "uint16", "Uint16", "u"},
		{"", "int32", "Int32", "i"},
		{"", "uint32", "Uint32", "u"},
		{"", "int64", "Int64", "i"},
		{"", "uint64", "Uint64", "u"},
		{"", "uintptr", "Uintptr", "u"},
		{"", "float32", "Float32", "f"},
		{"", "float64", "Float64", "f"},
		{"", "complex64", "Complex64", "c"},
		{"", "complex128", "Complex128", "c"},
		{"", "rune", "Rune", "r"},
		{"", "string", "String", "s"},
		{"time", "time.Duration", "Duration", "d"},
	}
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`

import (
{{- range .}}
	{{- if .PackagePath}}
	"{{.PackagePath}}"
	{{- end}}
{{- end}}
)
{{- range .}}

type {{.CapitalizedName}} struct {
       value {{.ReferenceName}}
       hasValue bool
}

func Make{{.CapitalizedName}}(value {{.ReferenceName}}) {{.CapitalizedName}} {
       return {{.CapitalizedName}}{value, true}
}

func ({{.ShortenName}} *{{.CapitalizedName}}) Set(value {{.ReferenceName}}) {
	*{{.ShortenName}} = Make{{.CapitalizedName}}(value)
}

func ({{.ShortenName}} *{{.CapitalizedName}}) Clear() {
	*{{.ShortenName}} = {{.CapitalizedName}}{}
}

func ({{.ShortenName}} {{.CapitalizedName}}) Value() {{.ReferenceName}} { return {{.ShortenName}}.value }
func ({{.ShortenName}} {{.CapitalizedName}}) HasValue() bool { return {{.ShortenName}}.hasValue }
{{- end}}
`)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
