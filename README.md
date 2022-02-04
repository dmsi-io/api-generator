# api-generator

This API generator was created as a result of the cumbersome process of creating Go structs equivalent to provided JSON responses. This initial version is able generate structs and a base GET request with or without pagination. However, the default usage has been customized specifically for responses for Agility and all data to create structs for must fall under the top level "response" key in the JSON object.

Additionally, this initial version only supports REST requests for GET HTTP methods.

### Usage

This repo has a provided `example.json` that shows the required inputs.

- The `response` object is where you will need to populate the example response from Agility.
- The top-level `method` string represents the name of the method and will be prefixed to certain variables and structs within the generated code.
- The top-level `jsonapi` array represents the data objects within the `response` that JSONAPI interface functions should be generated for.
- The top-level `endpoint` string represents the actual endpoint that is used to retrieve this data.

Once all data fields are updated in the JSON file the following command can be run locally from this repo:

```bash
go run main.go --file=<filename>
```

If your response uses the standard pagination to chunk data in the response an optional flag is available and can be used with:

```bash
go run main.go --file=<filename> --pagination
```

If running this locally within this repo, files matching the provided file input will be generated under the `rest` folder (`<filename>.go` and `<filename>_structs.go`). These can then be moved into your working repo.

Alternatively, this CLI tool can be installed as global command:

```bash
go install github.com/dmsi-io/api-generator@latest
```

And instead run it as follows in your repo of choice:

```bash
api-generator --file=<filename>
```

> Be weary that the JSON file should be at the top level of your directly and generated files will be added to the `rest` folder at the same level.
