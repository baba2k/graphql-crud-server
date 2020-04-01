package schema

import (
	"log"

	"github.com/baba2k/graphql-crud-server/storage"
	_type "github.com/baba2k/graphql-crud-server/type"
	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func CreateSchema(schema string, db storage.MongoDB) (graphql.Schema, error) {
	schem, gqlErr := gqlparser.LoadSchema(&ast.Source{Input: schema})
	if gqlErr != nil {
		log.Fatalf("can not load schema: " + gqlErr.Error())
	}

	schemaConfig := graphql.SchemaConfig{}

	// input
	input := map[string]graphql.Input{}
	for k, v := range _type.CreateInputObjects(schem) {
		input[k] = v
		input[k+"!"] = graphql.NewNonNull(v)
		input["["+k+"]"] = graphql.NewList(v)
		input["["+k+"!]"] = graphql.NewList(graphql.NewNonNull(v))
		input["["+k+"]!"] = graphql.NewNonNull(graphql.NewList(v))
		input["["+k+"!]!"] = graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(v)))
	}
	for k, v := range _type.ScalarTypes {
		input[k] = v
	}

	// output
	output := map[string]graphql.Output{}
	for k, v := range _type.CreateObjects(schem) {
		output[k] = v
		output[k+"!"] = graphql.NewNonNull(v)
		output["["+k+"]"] = graphql.NewList(v)
		output["["+k+"!]"] = graphql.NewList(graphql.NewNonNull(v))
		output["["+k+"]!"] = graphql.NewNonNull(graphql.NewList(v))
		output["["+k+"!]!"] = graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(v)))
	}
	for k, v := range _type.ScalarTypes {
		output[k] = v
	}

	schemaConfig.Query = _type.CreateQuery(schem, input, output, db)
	schemaConfig.Mutation = _type.CreateMutation(schem, input, output, db)
	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("can not create schema: " + err.Error())
	}
	return s, err
}
