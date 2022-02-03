package jenshared

import (
	"fmt"

	"github.com/alehechka/api-generator/utils"
	"github.com/dave/jennifer/jen"
)

// CreateJSONAPIEndpoint creates JSONAPI endpoint function
func CreateJSONAPIEndpoint(packageName, fileName string, m map[string]interface{}, pagination bool) {

	method := m["method"].(string)
	rootResponse := createRootResponse(method)
	endpoint := m["endpoint"].(string)
	endpointVar := createEndpoint(method)

	f := jen.NewFile(packageName)

	createJSONAPIEndpoint(f, fileName, rootResponse, endpointVar, endpoint, pagination)

	err := utils.CreateFilePath(packageName)
	utils.Check(err)

	err = f.Save(fmt.Sprintf("%s/%s.go", packageName, fileName))
	utils.Check(err)
}

func createJSONAPIEndpoint(f *jen.File, fileName, rootResponse, endpointVar, endpoint string, pagination bool) {
	addEndpoint(f, endpointVar, endpoint)
	addFuncHeader(f, fileName, rootResponse, endpointVar, pagination)
}

func addEndpoint(f *jen.File, endpointVar, endpoint string) {
	f.Const().Id(endpointVar).Op("=").Lit(endpoint)
}

func addFuncHeader(f *jen.File, fileName string, rootResponse, endpointVar string, pagination bool) {

	code := make([]jen.Code, 0)
	code = append(code, addLoggerDeclaration())
	code = append(code, addRequestURLDeclaration(endpointVar))
	code = append(code, addHTTPRequestDeclaration())
	code = append(code, jen.Line())
	if pagination {
		code = append(code, addPaginationHeaders()...)
		code = append(code, jen.Line())
		code = append(code, addPaginationRequestBody())
	} else {
		code = append(code, addEmptyRequestBody())
	}
	code = append(code, jen.Line())
	code = append(code, addPostRequest())

	code = append(code, jen.Line())
	code = append(code, addResponseBodyParsing(rootResponse)...)

	code = append(code, jen.Line())
	code = append(code, addDataArrayCreator()...)

	code = append(code, jen.Line())
	code = append(code, addErrors())

	code = append(code, jen.Line())
	if pagination {
		code = append(code, addPaginationLinks(fileName))
		code = append(code, jen.Line())
		code = append(code, addPaginationGinJSON())
	} else {
		code = append(code, addGinJSON())
	}

	f.Func().Id(fileName).Params(jen.Id("c").Op("*").Qual("github.com/gin-gonic/gin", "Context")).Block(code...)
}

func addLoggerDeclaration() *jen.Statement {
	return jen.Id("log").Op(":=").Qual("github.com/dmsi-io/go-utils/ginshared", "MustGetLogger").Params(jen.Id("c"))
}

func addRequestURLDeclaration(endpointVar string) *jen.Statement {
	return jen.Id("requestURL").Op(":=").Qual("github.com/dmsi-io/go-utils/ginshared", "MustCreateRequestURL").Params(jen.Id("c"), jen.Id(endpointVar))
}

func addHTTPRequestDeclaration() *jen.Statement {
	return jen.Id("httpRequest").Op(":=").Qual("github.com/dmsi-io/go-utils/ginshared", "MustGetRequestParams").Params(jen.Id("c"))
}

func addPaginationHeaders() []jen.Code {
	code := make([]jen.Code, 0)

	code = append(code, jen.List(jen.Id("chunkStartPointer"), jen.Id("_")).Op(":=").Qual("strconv", "Atoi").Params(jen.Qual("github.com/dmsi-io/go-utils/jsonapi", "GetQueryParameter").Params(jen.Id("c"), jen.Qual("github.com/dmsi-io/go-utils/jsonapi", "PageOffset"))))
	code = append(code, jen.List(jen.Id("recordFetchLimit"), jen.Id("_")).Op(":=").Qual("strconv", "Atoi").Params(jen.Qual("github.com/dmsi-io/go-utils/jsonapi", "GetQueryParameter").Params(jen.Id("c"), jen.Qual("github.com/dmsi-io/go-utils/jsonapi", "PageLimit"))))
	code = append(code, jen.Id("applyDataChunking").Op(":=").Id("recordFetchLimit").Op("!=").Lit(0))

	return code
}

func addEmptyRequestBody() *jen.Statement {
	return jen.Id("httpRequest").Dot("RequestBody").Op("=").Map(jen.String()).Interface().Block()
}

func addPaginationRequestBody() *jen.Statement {
	return jen.Id("httpRequest").Dot("RequestBody").Op("=").Map(jen.String()).Interface().Block(
		jen.Id("ApplyDataChunking").Op(":").Id("applyDataChunking").Op(","),
		jen.Id("ChunkStartPointer").Op(":").Id("chunkStartPointer").Op(","),
		jen.Id("RecordFetchLimit").Op(":").Id("recordFetchLimit").Op(","),
	)
}

func addPostRequest() *jen.Statement {
	return jen.List(jen.Id("response"), jen.Id("_")).Op(":=").Id("httpRequest").Dot("Post").Params(jen.Id("requestURL"))
}

func addResponseBodyParsing(rootResponse string) []jen.Code {
	code := make([]jen.Code, 0)

	code = append(code, jen.Var().Id("body").Id(rootResponse))
	code = append(code, jen.If(jen.Id("response").Op("!=").Nil()).Block(
		jen.Defer().Id("response").Dot("Body").Dot("Close").Params(),
		jen.Qual("encoding/json", "NewDecoder").Params(jen.Id("response").Dot("Body")).Dot("Decode").Params(jen.Op("&").Id("body")),
	))

	return code
}

func addDataArrayCreator() []jen.Code {
	code := make([]jen.Code, 0)

	code = append(code, jen.Id("data").Op(":=").Make(jen.Op("[]").Qual("github.com/dmsi-io/go-utils/jsonapi", "Data"), jen.Lit(0)))
	code = append(code, jen.For(jen.List(jen.Id("_"), jen.Id("value")).Op(":=").Range().Id("body").Dot("Response").Block(
		jen.Id("data").Op("=").Append(jen.Id("data"), jen.Id("value")),
	)))

	return code
}

func addErrors() *jen.Statement {
	return jen.Id("errs").Op(":=").Qual("github.com/dmsi-io/go-utils/jsonapi", "CreateErrorsFromResponse").Params(
		jen.Id("response").Dot("StatusCode"),
		jen.Id("body").Dot("Response").Dot("ReturnCode"),
		jen.Id("body").Dot("Response").Dot("MessageNum"),
		jen.Id("body").Dot("Response").Dot("MessageText"),
		jen.Id("log"),
	)
}

func addGinJSON() *jen.Statement {
	return jen.Id("c").Dot("JSON").Params(
		jen.Id("response").Dot("StatusCode"), jen.Qual("github.com/dmsi-io/go-utils/ginshared", "CreateJSONAPIResponse").Params(
			jen.Id("c"), jen.Id("data"), jen.Nil(), jen.Id("errs"), jen.Nil(), jen.Nil(),
		))
}

func addPaginationLinks(fileName string) *jen.Statement {
	return jen.Id("links").Op(":=").Qual("github.com/dmsi-io/go-utils/jsonapi", "CreateNextLinksFromPaginationResponse").Params(
		jen.Lit(fmt.Sprintf("/%s", fileName)),
		jen.Nil(),
		jen.Id("body").Dot("Response").Dot("MoreResultsAvailable"),
		jen.Id("body").Dot("Response").Dot("NextChunkStartPointer"),
		jen.Id("recordFetchLimit"),
	)
}

func addPaginationGinJSON() *jen.Statement {
	return jen.Id("c").Dot("JSON").Params(
		jen.Id("response").Dot("StatusCode"), jen.Qual("github.com/dmsi-io/go-utils/ginshared", "CreateJSONAPIResponse").Params(
			jen.Id("c"), jen.Id("data"), jen.Nil(), jen.Id("errs"), jen.Id("links"), jen.Nil(),
		))
}
