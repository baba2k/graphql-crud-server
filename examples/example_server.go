package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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

	opt := options.Client()
	opt = opt.SetHosts([]string{"localhost"})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongodb.NewMongoDB(ctx, opt, "graphql")
	if err != nil {
		panic("can not connect to database: " + err.Error())
	}

	q := (*schema).Query
	if q != nil {
		fields := graphql.Fields{}
		for _, field := range q.Fields {
			if strings.HasPrefix(field.Name, "_") {
				continue
			}
			name := field.Name
			log.Println("Found field: ", name)
			fields[name] = &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					return db.ReadAll(ctx, name)
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
