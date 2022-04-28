package main

import (
	"flag"

	"github.com/dmsi-io/api-generator/jenshared"
	"github.com/dmsi-io/api-generator/utils"
)

func main() {
	rest := flag.Bool("rest", false, "Flag to generate rest endpoint handler")
	gql := flag.Bool("gql", false, "Flag to generate gql query (Not currently supported)")
	fileName := flag.String("file", "", "File Name of JSON object")
	methodName := flag.String("method", "", "Method name of top level object")
	topLevelObject := flag.String("top-level", "response", "Top-level object to parse from JSON")
	pagination := flag.Bool("pagination", false, "Flag to add pagination params")
	jsonapiStructs := flag.String("jsonapi", "", "CSV list of object names to create JSON:API interfaces for.")
	endpoint := flag.String("endpoint", "", "Backend endpoint to send request to (only required for rest/gql)")
	flag.Parse()

	cleanFileName := utils.RemoveFileExtension(*fileName)
	fileName = &cleanFileName

	m, err := utils.GetJSON(*fileName)
	utils.Check(err)

	packageName := "rest"
	if *gql {
		packageName = "gql"
	}

	jenshared.CreateStructs(m, packageName, *fileName, *methodName, *topLevelObject, *jsonapiStructs)

	if *rest {
		jenshared.CreateJSONAPIEndpoint("rest", *fileName, m, *methodName, *endpoint, *topLevelObject, *pagination)
	}
}
