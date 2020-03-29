package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_type "github.com/baba2k/graphql-rungen/examples/type"
	"github.com/baba2k/graphql-rungen/storage/mongodb"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	/*
			parser.ParseString(
						`schema {
		    				query: CreateQuery
						}

						type CreateQuery {
							name: String!
						}`)
	*/

	// parse example graphql schema from file
	// graphql.ParseFile("examples/schema/university.graphql")

	content, err := ioutil.ReadFile("examples/schema/todo_simple.graphql")
	if err != nil {
		panic(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	schema := gqlparser.MustLoadSchema(&ast.Source{
		Name:  "todo_simple.graphql",
		Input: text,
	})
	schemaConfig := graphql.SchemaConfig{}

	opt := options.Client()
	opt = opt.SetHosts([]string{"localhost"})
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	db, err := mongodb.NewMongoDB(ctx, opt, "graphql")
	if err != nil {
		panic("can not connect to database: " + err.Error())
	}

	_type.CreateObjects(schema)
	schemaConfig.Query = _type.CreateQuery(schema, db)
	schemaConfig.Mutation = _type.CreateMutation(schema, db)

	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &s,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})
	http.Handle("/graphql", h)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("graphql server can not be started: " + err.Error())
	}
}
