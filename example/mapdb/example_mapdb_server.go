package main

import (
	"log"

	"github.com/baba2k/graphql-crud-server"
	"github.com/baba2k/graphql-crud-server/storage"
)

func main() {
	// load schema from file
	schema, err := graphql.LoadSchemaFromFile("example/mapdb/schema/todo.graphql")
	if err != nil {
		log.Fatal("can load schema: " + err.Error())
	}

	// create mapdb database interface
	db, err := storage.NewMapDB()
	if err != nil {
		log.Fatal("can not connect to database: " + err.Error())
	}

	// start server
	graphql.StartServer(":8080", schema, db)
}
