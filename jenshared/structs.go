package jenshared

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/dmsi-io/api-generator/cli"
	"github.com/dmsi-io/api-generator/utils"
)

type StructItem struct {
	JSONName string
	Name     string
	Type     string
}

type StructItems []StructItem

type StructItemMap map[string]StructItems

func CreateStructs(jsonMap map[string]interface{}, args cli.Arguments) {

	f := jen.NewFile(args.PackageName)

	AddStructsFromJSON(f, jsonMap, args.MethodName, args.TopLevelObject)
	GenerateJSONAPIInterfaceFunctions(f, args.JsonapiStructs)

	err := utils.CreateFilePath(args.PackageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s_structs.go", args.PackageName, args.FileName))
	utils.Check(err)
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
