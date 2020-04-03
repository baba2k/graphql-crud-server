package _type

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser/ast"
)

func CreateObjects(schema *ast.Schema, objects map[string]graphql.Output) map[string]graphql.Output {
	secondIteration := map[int]bool{}
	i := 0
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
			fName := f.Type.String()
			fields[f.Name] = &graphql.Field{
				Name:        f.Name,
				Type:        ScalarTypes[fName],
				Description: f.Description,
			}
			if fields[f.Name].Type == nil {
				fields[f.Name].Type = objects[fName]
				if fields[f.Name].Type == nil {
					secondIteration[i] = true
					break
				}
			}
		}

		if secondIteration[i] {
			i++
			continue
		}

		// object config
		objectConfig := graphql.ObjectConfig{
			Name:        pt[0].Name,
			Fields:      fields,
			Description: pt[0].Description,
		}

		k := objectConfig.Name
		v := graphql.NewObject(objectConfig)
		objects[k] = v
		objects[k+"!"] = graphql.NewNonNull(v)
		objects["["+k+"]"] = graphql.NewList(v)
		objects["["+k+"!]"] = graphql.NewList(graphql.NewNonNull(v))
		objects["["+k+"]!"] = graphql.NewNonNull(graphql.NewList(v))
		objects["["+k+"!]!"] = graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(v)))
		i++
	}

	for _, secondIter := range secondIteration {
		if secondIter {
			return CreateObjects(schema, objects)
		}
	}

	return objects
}
