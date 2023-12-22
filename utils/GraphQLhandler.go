package utils

import (
	"github.com/graphql-go/handler"
)

// GraphQLHandler handles GraphQL requests
var GraphQLHandler = handler.New(&handler.Config{
	Schema: &Schema,
	Pretty: true,
})
