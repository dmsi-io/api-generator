package main

import (
	"flag"

	"github.com/alehechka/api-generator/jenshared"
)

func main() {
	packageName := flag.String("package", "rest", "Package Name: options are rest or gql")
	fileName := flag.String("file", "", "File Name of JSON object (do not included file extension)")
	pagination := flag.Bool("pagination", false, "Flag to add pagination params")
	flag.Parse()

	if *packageName == "rest" {
		m := jenshared.CreateStructs(*packageName, *fileName)
		jenshared.CreateJSONAPIEndpoint(*packageName, *fileName, m, *pagination)
	}
}
