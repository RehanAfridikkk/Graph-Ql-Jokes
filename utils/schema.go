// handlers/schema.go
package utils

import (
	"github.com/graphql-go/graphql"
)

var jokeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Joke",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"joke": &graphql.Field{
				Type: jokeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetRandomJoke(), nil
				},
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)
