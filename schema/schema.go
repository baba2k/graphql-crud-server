package schema

import (
	"log"

	_type "github.com/baba2k/graphql-rungen/examples/type"
	"github.com/baba2k/graphql-rungen/storage/mongodb"
	"github.com/graphql-go/graphql"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func CreateSchema(schema string, db mongodb.MongoDB) (graphql.Schema, error) {
	schem, gqlErr := gqlparser.LoadSchema(&ast.Source{Input: schema})
	if gqlErr != nil {
		log.Fatalf("can not load schema: " + gqlErr.Error())
	}

	schemaConfig := graphql.SchemaConfig{}
	_type.CreateObjects(schem)
	schemaConfig.Query = _type.CreateQuery(schem, db)
	schemaConfig.Mutation = _type.CreateMutation(schem, db)
	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatal("can not create schema: " + err.Error())
	}
	return s, err
}
