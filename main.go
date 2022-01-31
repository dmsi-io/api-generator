package main

import (
	"fmt"

	"github.com/alehechka/api-generator/jenshared"
	"github.com/alehechka/api-generator/utils"
	. "github.com/dave/jennifer/jen"
)

func main() {
	m, err := utils.OpenJSONFile("branches.json")
	utils.Check(err)

	packageName := "rest"
	fileName := "data_structs"

	f := NewFile(packageName)

	jenshared.AddStructsFromJSON(f, m)
	jenshared.GenerateJSONAPIInterfaceFunctions(f, m["jsonapi"].([]interface{}))

	err = utils.CreateFilePath(packageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s.go", packageName, fileName))
	utils.Check(err)
}
