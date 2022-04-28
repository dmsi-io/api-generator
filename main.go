package main

import (
	"github.com/dmsi-io/api-generator/cli"
	"github.com/dmsi-io/api-generator/jenshared"
	"github.com/dmsi-io/api-generator/utils"
)

func main() {
	args := cli.InitializeCLIArguments()

	m, err := utils.GetJSON(args.FileName)
	utils.Check(err)

	jenshared.CreateStructs(m, args)

	if args.Rest {
		jenshared.CreateJSONAPIEndpoint(m, args)
	}
}
