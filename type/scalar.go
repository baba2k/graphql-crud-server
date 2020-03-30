package _type

import "github.com/graphql-go/graphql"

// scalar types String , Int , Float , Boolean , ID
var ScalarTypes = map[string]graphql.Type{
	"String":     graphql.String,
	"String!":    graphql.NewNonNull(graphql.String),
	"[String]":   graphql.NewList(graphql.String),
	"[String!]":  graphql.NewList(graphql.NewNonNull(graphql.String)),
	"[String]!":  graphql.NewNonNull(graphql.NewList(graphql.String)),
	"[String!]!": graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.String))),

	"Int":     graphql.Int,
	"Int!":    graphql.NewNonNull(graphql.Int),
	"[Int]":   graphql.NewList(graphql.Int),
	"[Int!]":  graphql.NewList(graphql.NewNonNull(graphql.Int)),
	"[Int]!":  graphql.NewNonNull(graphql.NewList(graphql.Int)),
	"[Int!]!": graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.Int))),

	"Float":     graphql.Float,
	"Float!":    graphql.NewNonNull(graphql.Float),
	"[Float]":   graphql.NewList(graphql.Float),
	"[Float!]":  graphql.NewList(graphql.NewNonNull(graphql.Float)),
	"[Float]!":  graphql.NewNonNull(graphql.NewList(graphql.Float)),
	"[Float!]!": graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.Float))),

	"Boolean":     graphql.Boolean,
	"Boolean!":    graphql.NewNonNull(graphql.Boolean),
	"[Boolean]":   graphql.NewList(graphql.Boolean),
	"[Boolean!]":  graphql.NewList(graphql.NewNonNull(graphql.Boolean)),
	"[Boolean]!":  graphql.NewNonNull(graphql.NewList(graphql.Boolean)),
	"[Boolean!]!": graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.Boolean))),

	"ID":     graphql.ID,
	"ID!":    graphql.NewNonNull(graphql.ID),
	"[ID]":   graphql.NewList(graphql.ID),
	"[ID!]":  graphql.NewList(graphql.NewNonNull(graphql.ID)),
	"[ID]!":  graphql.NewNonNull(graphql.NewList(graphql.ID)),
	"[ID!]!": graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.ID))),
}
