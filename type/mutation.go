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

func CreateMutation(schema *ast.Schema, input map[string]graphql.Input, output map[string]graphql.Output, db storage.MongoDB) *graphql.Object {
	m := (*schema).Mutation
	if m == nil {
		return nil
	}

	fields := graphql.Fields{}
	for _, field := range m.Fields {
		if strings.HasPrefix(field.Name, "_") {
			continue
		}
		name := field.Name
		log.Println("found mutation: " + name)

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
		if strings.HasPrefix(name, "create") {
			fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				collection := strings.ToLower(strings.TrimLeft(name, "create"))
				res, err := db.Create(ctx, collection, p.Args)
				if err != nil {
					return nil, err
				}
				return res, err
			}
		} else if strings.HasPrefix(name, "update") {
			fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				res, err := db.Update(ctx, strings.ToLower(strings.TrimLeft(name, "update")), p.Args["id"], p.Args)
				if err != nil {
					return nil, err
				}
				return res, err
			}
		} else if strings.HasPrefix(name, "delete") {
			fields[name].Resolve = func(p graphql.ResolveParams) (interface{}, error) {
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				res, err := db.Delete(ctx, strings.ToLower(strings.TrimLeft(name, "delete")), p.Args["id"])
				if err != nil {
					return nil, err
				}
				return res, err
			}
		} else {
			log.Println("unknown field " + name)
		}
	}

	objectConfig := graphql.ObjectConfig{Name: m.Name, Fields: fields}

	return graphql.NewObject(objectConfig)
}
