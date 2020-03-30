package main

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/baba2k/graphql-rungen/server"
	"github.com/baba2k/graphql-rungen/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load schema from file
	schema, err := ioutil.ReadFile("examples/schema/todo_simple.graphql")
	if err != nil {
		panic(err)
	}

	// set database to mongodb
	opt := options.Client()
	opt = opt.SetHosts([]string{"localhost"})
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	db, err := mongodb.NewMongoDB(ctx, opt, "graphql")
	if err != nil {
		panic("can not connect to database: " + err.Error())
	}

	// start server
	server.StartServer(":8080", string(schema), db)
}
