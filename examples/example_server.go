package main

import (
	"github.com/baba2k/graphql-rungen"
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
	graphql.ParseFile("examples/schema/university.graphql")

	// TODO: Get objects by reflection

	// TODO: Set standard resolver interface to mongodb

	// TODO: Overwrite a resolver

	// TODO: Start graphql server
}
