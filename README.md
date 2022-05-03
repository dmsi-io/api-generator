# api-generator

This API generator was created as a result of the cumbersome process of creating Go structs equivalent to provided JSON responses. This initial version is able generate structs and a base GET request with or without pagination. However, the default usage has been customized specifically for responses for Agility and all data to create structs for must fall under the top level "response" key in the JSON object.

Additionally, this initial version only supports REST requests for GET HTTP methods.

### Installation

```bash
go install github.com/dmsi-io/api-generator@latest
```

### Usage

This repo has a provided `example.json` that displays some basic data that can be parsed into structs.

```bash
api-generator --file=example.json
```

This will result in a new file created under `rest/example_structs.go` with all Go structs created from the provided JSON.

Additionally, this CLI tool supports STDIN via piping a JSON body into the command:

```bash
cat example.json | api-generator --file=test

pbpaste | api-generator --file=test
```

> In this example the `file` CLI arg is still required to determine output file (i.e. `test_structs.go`)

### CLI Arguments

| Argument         | Example                              | Type         | Purpose                                                                                                                  |
| ---------------- | ------------------------------------ | ------------ | ------------------------------------------------------------------------------------------------------------------------ |
| REST             | `--rest`                             | `boolean`    | Will generate a Restful JSON:API GET request handler with `gin` framework                                                |
| GraphQL          | `--gql`                              | `boolean`    | (NOT FULLY SUPPORTED): Use to change package name of structs to `gql`. Future use will generate a GraphQL Query Handler. |
| Filename         | `--file=example.json --file=example` | `string`     | Attempts to open JSON file of same name. Output Go files will be named the same.                                         |
| Method Name      | `--method=Example`                   | `string`     | Applies the method name as a prefix to structs, variables, and functions.                                                |
| Top-level Object | `--top-level=request`                | `string`     | Selects the top-level object to start parsing from in JSON. Default is `response`                                        |
| Pagination       | `--pagination`                       | `boolean`    | Flag to add pagination logic to REST endpoint.                                                                           |
| JSON:API         | `--jsonapi=User,Customer`            | `string:csv` | CSV list of object names to generate JSON:API interface functions for.                                                   |
| Endpoint         | `--endpoint=example/path`            | `string`     | Backend endpoint path to make requests REST requests to.                                                                 |

### Local Development

While developing locally, the CLI can be used directly with the `go run` option as follows:

```bash
go run main.go --file=example.json <...args>
```
