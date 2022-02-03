package jenshared

import (
	"fmt"
	"strings"

	"github.com/alehechka/api-generator/utils"
	"github.com/dave/jennifer/jen"
)

type StructItem struct {
	JSONName string
	Name     string
	Type     string
}

type StructItems []StructItem

type StructItemMap map[string]StructItems

func CreateStructs(packageName, fileName string) map[string]interface{} {
	m, err := utils.OpenJSONFile(fmt.Sprintf("%s.json", fileName))
	utils.Check(err)

	f := jen.NewFile(packageName)

	AddStructsFromJSON(f, m)
	GenerateJSONAPIInterfaceFunctions(f, m["jsonapi"].([]interface{}))

	err = utils.CreateFilePath(packageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s_structs.go", packageName, fileName))
	utils.Check(err)

	return m
}

func AddStructs(f *jen.File, itemMap StructItemMap) {
	for name, items := range itemMap {
		f.Add(CreateStruct(name, items))
		f.Line()
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
