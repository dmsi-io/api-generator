package jenshared

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

type StructItem struct {
	JSONName string
	Name     string
	Type     string
}

type StructItems []StructItem

type StructItemMap map[string]StructItems

func AddStructs(f *jen.File, itemMap StructItemMap) {
	for name, items := range itemMap {
		f.Add(CreateStruct(name, items))
	}
}

func AddStruct(f *jen.File, name string, items StructItems) {
	f.Add(CreateStruct(name, items))
}

func CreateStruct(name string, items StructItems) *jen.Statement {

	structItems := CreateStructItems(items)

	return jen.Type().Id(name).Struct(structItems...)
}

func CreateStructItems(items StructItems) []jen.Code {
	structItems := make([]jen.Code, 0)

	for _, item := range items {
		structItems = append(structItems, CreateStructItem(item))
	}

	return structItems
}

func CreateStructItem(item StructItem) jen.Code {
	s := jen.Id(strings.Title(item.Name)).Id(item.Type)

	if item.JSONName != "" {
		s.Tag(map[string]string{"json": item.JSONName})
	}
	return s
}
