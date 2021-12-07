package resolvers

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/schartey/dgraph-lambda-example/lambda/model"
	"github.com/schartey/dgraph-lambda-go/api"
)

type MutationResolverInterface interface {
	Mutation_newUserGraphql(ctx context.Context, name string, authHeader api.AuthHeader) (string, *api.LambdaError)
}

type MutationResolver struct {
	*Resolver
}

func (q *MutationResolver) Mutation_newUserGraphql(ctx context.Context, name string, authHeader api.AuthHeader) (string, *api.LambdaError) {
	var respData struct {
		AddUser struct {
			Users []*model.User `json:"user"`
		} `json:"addUser"`
	}

	q.Gql.Run(ctx, graphql.NewRequest(fmt.Sprintf(`mutation {
		addUser(input: { username: "%s", email: "%s@mail.com", password: "123456" }) {
			user {
				id
				username
				email
			}
	   }
	 }`, name, name)), &respData)

	if len(respData.AddUser.Users) > 0 {
		return respData.AddUser.Users[0].Id, nil
	}
	return "", nil
}
