package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func main() {
	/*
			parser.ParseString(
						`schema {
		    				query: Query
						}

						type Query {
							name: String!
						}`)
	*/

	// parse example graphql schema from file
	// graphql.ParseFile("examples/schema/university.graphql")

	content, err := ioutil.ReadFile("examples/schema/hello.graphql")
	if err != nil {
		panic(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	schema := gqlparser.MustLoadSchema(&ast.Source{
		Name:  "hello.graphql",
		Input: text,
	})
	schemaConfig := graphql.SchemaConfig{}

	q := (*schema).Query
	if q != nil {
		fields := graphql.Fields{}
		for _, field := range q.Fields {
			fields[field.Name] = &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			}
		}
		rootQuery := graphql.ObjectConfig{Name: q.Name, Fields: fields}
		schemaConfig.Query = graphql.NewObject(rootQuery)
	}

	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// TODO: Get objects by reflection

	// TODO: Set standard resolver interface to mongodb

	// TODO: Start graphql server
	h := handler.New(&handler.Config{
		Schema:   &s,
		Pretty:   true,
		GraphiQL: false,
	})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
