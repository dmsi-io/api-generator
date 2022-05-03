package cli

import (
	"flag"
	"strings"

	"github.com/dmsi-io/api-generator/utils"
)

type Arguments struct {
	Rest           bool
	GraphQL        bool
	PackageName    string
	FileName       string
	MethodName     string
	TopLevelObject string
	Pagination     bool
	JsonapiStructs []string
	Endpoint       string
}

func InitializeCLIArguments() Arguments {
	rest := flag.Bool("rest", false, "Flag to generate rest endpoint handler")
	gql := flag.Bool("gql", false, "Flag to generate gql query (Not currently supported)")
	fileName := flag.String("file", "", "File Name of JSON object")
	methodName := flag.String("method", "", "Method name of top level object")
	topLevelObject := flag.String("top-level", "response", "Top-level object to parse from JSON")
	pagination := flag.Bool("pagination", false, "Flag to add pagination params")
	jsonapiCSV := flag.String("jsonapi", "", "CSV list of object names to create JSON:API interfaces for.")
	endpoint := flag.String("endpoint", "", "Backend endpoint to send request to (only required for rest/gql)")
	flag.Parse()

	var jsonapiStructs []string
	if len(*jsonapiCSV) > 0 {
		jsonapiStructs = strings.Split(*jsonapiCSV, ",")
	}

	packageName := "rest"
	if *gql {
		packageName = "gql"
	}

	return Arguments{
		Rest:           *rest,
		GraphQL:        *gql,
		PackageName:    packageName,
		FileName:       utils.RemoveFileExtension(*fileName),
		MethodName:     *methodName,
		TopLevelObject: *topLevelObject,
		Pagination:     *pagination,
		JsonapiStructs: jsonapiStructs,
		Endpoint:       *endpoint,
	}
}
