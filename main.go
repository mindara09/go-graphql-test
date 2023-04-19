package main

import (
	"encoding/json"
	"fmt"
	"go-graphql-test/authentication"
	"go-graphql-test/schema/mutation"
	"go-graphql-test/schema/query"
	"net/http"

	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	var Schame, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.RootQuery,
		Mutation: mutation.Mutation,
	})
	if err != nil {
		panic(err.Error())
	}
	// 1.Route GraphQL
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), Schame)
		json.NewEncoder(w).Encode(result)
	})
	// 2.Route Login
	http.HandleFunc("/login", authentication.CreateTokenEndpoint)

	// Serve
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
