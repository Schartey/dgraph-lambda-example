package resolvers

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/machinebox/graphql"
)

// Add objects to your desire
type Resolver struct {
	Gql *graphql.Client
	Dql *dgo.Dgraph
}
