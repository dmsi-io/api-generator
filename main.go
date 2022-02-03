package main

import (
	"os"

	"github.com/alehechka/api-generator/jenshared"
)

func main() {
	packageName := os.Args[1]
	fileName := os.Args[2]

	var page bool
	if len(os.Args) > 3 {
		page = os.Args[3] == "true"
	}

	m := jenshared.CreateStructs(packageName, fileName)

	jenshared.CreateJSONAPIEndpoint(packageName, fileName, m, page)
}
