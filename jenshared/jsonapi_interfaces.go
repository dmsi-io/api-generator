package jenshared

import (
	"strings"

	"github.com/dave/jennifer/jen"
)

// GenerateJSONAPIInterfaceFunctions generates jsonapi.Data interface functions
func GenerateJSONAPIInterfaceFunctions(f *jen.File, structNames []interface{}) {
	for _, structName := range structNames {
		name := structName.(string)
		f.Add(generateIDFunc(name))
		f.Add(generateTypeFunc(name))
		f.Add(generateAttributesFunc(name))
		f.Add(generateLinksFunc(name))
		f.Add(generateRelationshipsFunc(name))
		f.Add(generateMetaFunc(name))
	}
}

func generateIDFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("ID").Params().String().Block(jen.Return(jen.Id("d").Dot("ID")))

	return c
}

func generateTypeFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("Type").Params().String().Block(jen.Return(jen.Id("\"\"")))

	return c
}

func generateAttributesFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("Attributes").Params().Interface().Block(jen.Return(jen.Id("d")))

	return c
}

func generateLinksFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("Links").Params().Qual("github.com/dmsi-io/go-utils/jsonapi", "Links").Block(jen.Return(jen.Qual("github.com/dmsi-io/go-utils/jsonapi", "Links").Block()))

	return c
}

func generateRelationshipsFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("Relationships").Params().Map(jen.String()).Qual("github.com/dmsi-io/go-utils/jsonapi", "Relationship").Block(jen.Return(jen.Nil()))

	return c
}

func generateMetaFunc(structName string) *jen.Statement {

	c := jen.Func().Params(
		jen.Id("d").Id(strings.Title(structName)),
	).Id("Meta").Params().Interface().Block(jen.Return(jen.Nil()))

	return c
}
