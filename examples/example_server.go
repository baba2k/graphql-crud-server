package main

import (
	"context"
	"time"

	"github.com/baba2k/graphql-rungen"
	"github.com/baba2k/graphql-rungen/storage"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load schema from file
	err, schema := graphql.LoadSchemaFromFile("examples/schema/todo_simple.graphql")

	// create mongodb interface
	opt := options.Client()
	opt = opt.SetHosts([]string{"localhost"})
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	db, err := storage.NewMongoDB(ctx, opt, "graphql")
	if err != nil {
		panic("can not connect to database: " + err.Error())
	}

	// start server
	graphql.StartServer(":8080", schema, db)
}
