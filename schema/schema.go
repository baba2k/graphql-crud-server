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
	for k, v := range _type.ScalarTypes {
		input[k] = v
	}
	input = _type.CreateInputObjects(schem, input)

	// output
	output := map[string]graphql.Output{}
	for k, v := range _type.ScalarTypes {
		output[k] = v
	}
	output = _type.CreateObjects(schem, output)

	schemaConfig.Query = _type.CreateQuery(schem, input, output, db)
	schemaConfig.Mutation = _type.CreateMutation(schem, input, output, db)
	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("can not create schema: " + err.Error())
	}
	return s, err
}
