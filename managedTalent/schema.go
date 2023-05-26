package managedTalent

import (
	"api/di"
	"github.com/graphql-go/graphql"
	"log"
)

func InitializeQuery(query graphql.Fields) graphql.Fields {
	var hubberResolver *resolver
	err := di.Container.Invoke(func(resolver *resolver) {
		hubberResolver = resolver
	})

	if err != nil {
		log.Fatalf("failed initialization hubber resolver: %v", err)
	}

	query["hubbers"] = &graphql.Field{
		Type:    graphql.NewList(hubberType),
		Resolve: hubberResolver.hubbers,
	}

	return query
}

type Hubber struct {
	Id   string
	Code string
	Name string
}

var hubberType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Hubber",
	Description: "a Hubber",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if hubber, ok := p.Source.(Hubber); ok {
					return hubber.Id, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if hubber, ok := p.Source.(Hubber); ok {
					return hubber.Name, nil
				}
				return nil, nil
			},
		},
		"code": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if hubber, ok := p.Source.(Hubber); ok {
					return hubber.Code, nil
				}
				return nil, nil
			},
		},
	},
})
