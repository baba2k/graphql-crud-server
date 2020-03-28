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

	// query
	q := (*schema).Query
	if q != nil {
		fields := graphql.Fields{}
		for _, field := range q.Fields {
			if strings.HasPrefix(field.Name, "_") {
				continue
			}
			name := field.Name
			log.Println("Found query field: " + name)

			// set field type
			fields[name] = &graphql.Field{
				Type: scalarTypesMap[field.Type.String()],
			}

			// set field arguments
			fields[name].Args = make(graphql.FieldConfigArgument)
			for _, arg := range field.Arguments {
				fields[name].Args[arg.Name] = &graphql.ArgumentConfig{
					Type:         scalarTypesMap[arg.Type.String()],
					DefaultValue: arg.DefaultValue,
					Description:  arg.Description,
				}
			}

			// set field resolver
			if strings.HasSuffix(name, "s") {
				fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					return db.ReadAll(ctx, strings.ToLower(strings.TrimRight(strings.TrimLeft(name, "read"), "s")))
				}
			} else {
				fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					return db.ReadOne(ctx, strings.ToLower(strings.TrimLeft(name, "read")), p.Args["id"])
				}
			}
		}
		rootQuery := graphql.ObjectConfig{Name: q.Name, Fields: fields}
		schemaConfig.Query = graphql.NewObject(rootQuery)
	}

	// mutation
	m := (*schema).Mutation
	if m != nil {
		fields := graphql.Fields{}
		for _, field := range m.Fields {
			if strings.HasPrefix(field.Name, "_") {
				continue
			}
			name := field.Name
			log.Println("Found mutation field: " + name)

			// set field type
			fields[name] = &graphql.Field{
				Type: scalarTypesMap[field.Type.String()],
			}

			// set field arguments
			fields[name].Args = make(graphql.FieldConfigArgument)
			for _, arg := range field.Arguments {
				fields[name].Args[arg.Name] = &graphql.ArgumentConfig{
					Type:         scalarTypesMap[arg.Type.String()],
					DefaultValue: arg.DefaultValue,
					Description:  arg.Description,
				}
			}

			// set field resolver
			if strings.HasPrefix(name, "create") {
				fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					return db.Create(ctx, strings.ToLower(strings.TrimLeft(name, "create")), p.Args)
				}
			} else if strings.HasPrefix(name, "update") {
				fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					err := db.Update(ctx, strings.ToLower(strings.TrimLeft(name, "update")), p.Args, p.Args["id"])
					return nil, err
				}
			} else if strings.HasPrefix(name, "delete") {
				fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
					ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
					err := db.Delete(ctx, strings.ToLower(strings.TrimLeft(name, "delete")), p.Args["id"])
					return nil, err
				}
			} else {
				log.Println("unknown field " + name)
			}
		}
		rootMutation := graphql.ObjectConfig{Name: m.Name, Fields: fields}
		schemaConfig.Mutation = graphql.NewObject(rootMutation)
	}

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

// scalar types String , Int , Float , Boolean , ID
var scalarTypesMap = map[string]graphql.Type{
	"String!":  graphql.NewNonNull(graphql.String),
	"String":   graphql.String,
	"[String]": graphql.NewList(graphql.String),
	"Int!":     graphql.NewNonNull(graphql.Int),
	"Int":      graphql.Int,
	"Float!":   graphql.NewNonNull(graphql.Float),
	"Float":    graphql.Float,
	"Boolean!": graphql.NewNonNull(graphql.Boolean),
	"Boolean":  graphql.Boolean,
	"ID!":      graphql.NewNonNull(graphql.ID),
	"ID":       graphql.ID,
}
