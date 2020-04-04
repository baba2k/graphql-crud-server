package main

import (
	"context"
	"log"
	"time"

	"github.com/baba2k/graphql-crud-server"
	"github.com/baba2k/graphql-crud-server/storage"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load schema from file
	schema, err := graphql.LoadSchemaFromFile("example/schema/todo.graphql")
	if err != nil {
		log.Fatal("can load schema: " + err.Error())
	}

	// create mongodb database interface
	opt := options.Client()
	opt = opt.SetHosts([]string{"localhost"})
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	db, err := storage.NewMongoDB(ctx, opt, "graphql")
	if err != nil {
		log.Fatal("can not connect to database: " + err.Error())
	}

	// start server
	graphql.StartServer(":8080", schema, db)
}
