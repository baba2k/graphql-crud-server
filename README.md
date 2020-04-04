# graphql-crud-server

GraphQL CRUD server library for golang

## What is graphql-crud-server?

graphql-crud-server is a GraphQL library written in golang for building generic CRUD (Create, Read, Update, Delete) graphql servers fast and simple. No need for implementing resolvers, models or other overhead. Just load a graphql schema and start the server. See section "Examples" for more details.

## Installation

```bash
# Go Modules
require github.com/baba2k/graphql-crud-server v0.1.5
```

## Usage

The following samples will assist you to become as comfortable as possible with graphql-crud-server library.

```go
// Import graphql-crud-server into your code and refer it as `graphql`.
import "github.com/baba2k/graphql-crud-server"
```
## Schema definition

### Query

* Name of the collection / table  will be parsed from the field definition in lowercase
* *"read"* prefix and *"s"* suffix will be cut off
* List as return type defines a field which returns a list
* All other fields return an object

#### Example
```graphql
type Query {
    readTodos: [Todo]
    readTodo(id: ID!): Todo
}
```
* The collection / table is *"todo"*
* *readTodos* returns all todos
* *readTodo* returns a todo by id
* `todo(id: ID!): Todo` and `todos: [Todo]` would be also fine

### Mutation

TODO

## Supported database interfaces

### MongoDB

Create a MongoDB interface with the official mongo driver (https://github.com/mongodb/mongo-go-driver) and pass it to the graphql-crud-server as argument. See section "Examples" for more details.

## Examples

[Simple "Todo" example](./example/example_todo_server.go)

See "example" dir for more examples

## License

[MIT](LICENSE)