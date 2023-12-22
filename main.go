package main

import (
	"net/http"

	"Graph-QL-Jokes/utils"
)

func main() {
	utils.SetupDB()

	handlerWithMiddleware := utils.Middleware(utils.GraphQLHandler)

	http.Handle("/graphql", handlerWithMiddleware)
	http.ListenAndServe(":8080", nil)
}
