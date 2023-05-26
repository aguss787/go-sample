package graphqlServer

import (
	"github.com/graphql-go/graphql"
)

type InitializerFunc func(query graphql.Fields) graphql.Fields

func buildSchema(queryInitializer []InitializerFunc) (graphql.Schema, error) {
	fields := graphql.Fields{}

	for _, f := range queryInitializer {
		fields = f(fields)
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}
