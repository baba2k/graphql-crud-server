package _type

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/baba2k/graphql-crud-server/storage"
	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func CreateQuery(schema *ast.Schema, input map[string]graphql.Input, output map[string]graphql.Output, db storage.MongoDB) *graphql.Object {
	q := (*schema).Query
	if q == nil {
		return nil
	}

	fields := graphql.Fields{}
	for _, field := range q.Fields {
		if strings.HasPrefix(field.Name, "_") {
			continue
		}
		name := field.Name
		log.Println("found query: " + name)

		// set field type
		fields[name] = &graphql.Field{
			Type: output[field.Type.String()],
		}

		// set field arguments
		fields[name].Args = make(graphql.FieldConfigArgument)
		for _, arg := range field.Arguments {
			fields[name].Args[arg.Name] = &graphql.ArgumentConfig{
				Type:         input[arg.Type.String()],
				DefaultValue: arg.DefaultValue,
				Description:  arg.Description,
			}
		}

		// set field resolver
		if strings.HasPrefix(field.Type.String(), "[") {
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

	objectConfig := graphql.ObjectConfig{Name: q.Name, Fields: fields}

	return graphql.NewObject(objectConfig)
}
