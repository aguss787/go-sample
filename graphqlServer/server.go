package graphqlServer

import (
	"github.com/graphql-go/handler"
	"net/http"
)

func RunServer(queryInitializer []InitializerFunc) error {
	schema, err := buildSchema(queryInitializer)
	if err != nil {
		return err
	}

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", handler)
	return http.ListenAndServe(":8080", nil)
}
