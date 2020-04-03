package _type

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func CreateObjects(schema *ast.Schema) map[string]*graphql.Object {
	objects := map[string]*graphql.Object{}

	for _, pt := range schema.PossibleTypes {
		if len(pt) != 1 {
			log.Println("unknown error on while creating objects: len(possibleType) != 1")
			continue
		}
		if pt[0].BuiltIn || pt[0].Name == "Query" || pt[0].Name == "Mutation" || pt[0].IsInputType() {
			continue
		}

		log.Println("found object: " + pt[0].Name)

		// fields
		fields := graphql.Fields{}
		for _, f := range pt[0].Fields {
			fields[f.Name] = &graphql.Field{
				Name:        f.Name,
				Type:        ScalarTypes[f.Type.String()],
				Description: f.Description,
			}
		}

		// object config
		objectConfig := graphql.ObjectConfig{
			Name:        pt[0].Name,
			Fields:      fields,
			Description: pt[0].Description,
		}

		objects[objectConfig.Name] = graphql.NewObject(objectConfig)
	}

	return objects
}
