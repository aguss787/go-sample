package managedTalent

import "github.com/graphql-go/graphql"

type resolver struct {
	hubberApi HubberAPI
}

func (r *resolver) hubbers(p graphql.ResolveParams) (interface{}, error) {
	return r.hubberApi.GetAllHubbers()
}
