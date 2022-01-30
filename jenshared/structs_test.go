package jenshared

import (
	"fmt"
	"testing"

	"github.com/alehechka/api-generator/utils"
	. "github.com/dave/jennifer/jen"
)

func TestAddStruct(t *testing.T) {
	packageName := "test_output"
	fileName := "output"

	f := NewFile(packageName)

	items := StructItems{
		{
			Name: "waterWay",
			Type: "int",
		},
		{
			Name: "xRay",
			Type: "float64",
		},
		{
			Name: "yamlParse",
			Type: "bool",
		},
		{
			Name: "zebraRun",
			Type: "string",
		},
		{
			Name: "customStruct",
			Type: "CustomStruct",
		},
	}

	customItems := StructItems{{Name: "thing", Type: "int"}}

	itemMap := StructItemMap{"Foo": items, "CustomStruct": customItems}

	AddStructs(f, itemMap)

	err := utils.CreateFilePath(packageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s.go", packageName, fileName))
	utils.Check(err)
}
