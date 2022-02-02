package jenshared

import (
	"fmt"

	"github.com/alehechka/api-generator/utils"
	"github.com/dave/jennifer/jen"
)

func CreateJSONAPIEndpoint(packageName, fileName string) {
	f := jen.NewFile(packageName)

	err := utils.CreateFilePath(packageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s.go", packageName, fileName))
	utils.Check(err)
}
