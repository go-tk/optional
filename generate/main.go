//go:generate go run $GOFILE

package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"strings"
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
	typeInfos := makeTypeInfos()
	var buffer bytes.Buffer
	if err := template.Must(template.New("").Parse(`// Code generated. DO NOT EDIT.

package `+packageName+`
 {{- range .}}

 type {{.CapitalizedName}} struct {
	value {{.Name}}
	hasValue bool
 }

 func Make{{.CapitalizedName}}(value {{.Name}}) {{.CapitalizedName}} {
	return {{.CapitalizedName}}{value, true}
 }

 func ({{.ShortenName}} {{.CapitalizedName}}) Value() {{.Name}} { return {{.ShortenName}}.value }
 func ({{.ShortenName}} {{.CapitalizedName}}) HasValue() bool { return {{.ShortenName}}.hasValue }
 {{- end}}
 `)).Execute(&buffer, typeInfos); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

type typeInfo struct {
	Name            string
	CapitalizedName string
	ShortenName     string
}

func makeTypeInfos() []typeInfo {
	typeNames := []string{
		"bool",
		"byte",
		"int", "uint",
		"int8", "uint8",
		"int16", "uint16",
		"int32", "uint32",
		"int64", "uint64",
		"uintptr",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"rune",
		"string",
	}
	typeInfos := make([]typeInfo, len(typeNames))
	for i, typeName := range typeNames {
		typeInfo := &typeInfos[i]
		typeInfo.Name = typeName
		typeInfo.CapitalizedName = strings.ToUpper(typeName[:1]) + typeName[1:]
		typeInfo.ShortenName = typeName[:1]
	}
	return typeInfos
}
