# graphql-crud-server

GraphQL CRUD server library for golang

## What is graphql-crud-server?

graphql-crud-server is a GraphQL library written in golang for building generic CRUD (Create, Read, Update, Delete) graphql servers fast and simple. No need for implementing resolvers, models or other overhead. Just load a graphql schema and start the server. See sections "Schema definition" and "Examples" for more details.

## Installation

```bash
# go modules
require github.com/baba2k/graphql-crud-server v0.1.5
```

## Usage

The following samples will assist you to become as comfortable as possible with graphql-crud-server library.

```go
// import graphql-crud-server into your code and refer it as `graphql`
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

See section "Examples" for more examples.

### Mutation

* Name of the collection / table  will be parsed from the field definition in lowercase
* All fields must have one of these prefixes: *"create", "update" and "delete"* 
* Prefixes will be cut off
* All fields return the modified object

#### Example
```graphql
type Mutation {
    createTodo(todo: TodoInput!): Todo
    updateTodo(id: ID!, todo: TodoInput!): Todo
    deleteTodo(id: ID!): Todo
}
```
* The collection / table is "todo"

See section "Examples" for more examples.

## Supported database interfaces

### MapDB (in-memory)

Create a MapDB interface and pass it to the graphql-crud-server as argument. See section "Examples" for more details.

#### Example

```go
// create mapdb database interface
db, err := storage.NewMapDB()
if err != nil {
    log.Fatal("can not connect to database: " + err.Error())
}
```
### MongoDB

Create a MongoDB interface with the official mongo driver (https://github.com/mongodb/mongo-go-driver) and pass it to the graphql-crud-server as argument. See section "Examples" for more details.

#### Example

```go
// create mongodb database interface
opt := options.Client()
opt = opt.SetHosts([]string{"localhost"})
ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
db, err := storage.NewMongoDB(ctx, opt, "graphql")
if err != nil {
    log.Fatal("can not connect to database: " + err.Error())
}
```

## Examples

* [MapDB "Todo" example](./example/mapdb/example_mapdb_server.go)
* [MongoDB "Todo" example](./example/mongodb/example_mongodb_server.go)

See "example" dir for more examples

## License

[MIT](LICENSE)