package _type

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func CreateInputObjects(schema *ast.Schema) map[string]*graphql.InputObject {
	objects := map[string]*graphql.InputObject{}

	for _, pt := range schema.PossibleTypes {
		if len(pt) != 1 {
			log.Println("unknown error on while creating objects: len(possibleType) != 1")
			continue
		}
		if pt[0].BuiltIn || pt[0].Name == "Query" || pt[0].Name == "Mutation" || !pt[0].IsInputType() {
			continue
		}

		log.Println("found input: " + pt[0].Name)

		// fields
		fields := graphql.InputObjectConfigFieldMap{}
		for _, f := range pt[0].Fields {
			fields[f.Name] = &graphql.InputObjectFieldConfig{
				Type:        ScalarTypes[f.Type.String()],
				Description: f.Description,
			}
		}

		// object config
		objectConfig := graphql.InputObjectConfig{
			Name:        pt[0].Name,
			Fields:      fields,
			Description: pt[0].Description,
		}

		objects[objectConfig.Name] = graphql.NewInputObject(objectConfig)
	}

	return objects
}
