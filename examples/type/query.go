package _type

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/baba2k/graphql-rungen/storage/mongodb"
	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func CreateQuery(schema *ast.Schema, db mongodb.MongoDB) *graphql.Object {
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
		log.Println("Found query field: " + name)

		// set field type
		fields[name] = &graphql.Field{
			Type: ScalarTypes[field.Type.String()],
		}

		// set field arguments
		fields[name].Args = make(graphql.FieldConfigArgument)
		for _, arg := range field.Arguments {
			fields[name].Args[arg.Name] = &graphql.ArgumentConfig{
				Type:         ScalarTypes[arg.Type.String()],
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

	objectConfig := graphql.ObjectConfig{Name: q.Name, Fields: fields}

	return graphql.NewObject(objectConfig)
}
