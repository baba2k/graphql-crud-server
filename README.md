# graphql-crud-server

GraphQL generic CRUD server for golang

## What is graphql-crud-server?

graphql-crud-server is a GraphQL library written in golang for building generic CRUD graphql 
servers fast and simple. No need for any models or other overhead. Just load a graphql schema 
and start the server. See section "Examples" for more details.

## Installation

```bash
# Go Modules
require github.com/baba2k/graphql-crud-server v0.1.3
```

## Usage

The following samples will assist you to become as comfortable as possible 
with graphql-crud-server library.

```go
// Import graphql-crud-server into your code and refer it as `graphql`.
import "github.com/baba2k/graphql-crud-server"
```

## Examples

[Simple "Todo" example](./example/example_todo_server.go)

See "example" dir for more examples

## License

[MIT](LICENSE)