package resolvers

import (
	"context"
	"net/http"

	"github.com/schartey/dgraph-lambda-example/lambda/model"
	"github.com/schartey/dgraph-lambda-go/api"
	"github.com/schartey/dgraph-lambda-go/dson"
)

type QueryResolverInterface interface {
	Query_firstUserDql(ctx context.Context, authHeader api.AuthHeader) (*model.User, *api.LambdaError)
}

type QueryResolver struct {
	*Resolver
}

func (q *QueryResolver) Query_firstUserDql(ctx context.Context, authHeader api.AuthHeader) (*model.User, *api.LambdaError) {
	res, err := q.Dql.NewReadOnlyTxn().Query(ctx, `query {
		findUsers(func: type(User)) {
			uid
			User.username
			User.email
		}
	}`)
	if err != nil {
		return nil, &api.LambdaError{Underlying: err, Status: http.StatusInternalServerError}
	}

	var respData struct {
		Users []*model.User `dql:"findUsers"`
	}

	err = dson.Unmarshal(res.GetJson(), &respData)
	if err != nil {
		return nil, &api.LambdaError{Underlying: err, Status: http.StatusInternalServerError}
	}

	if len(respData.Users) == 0 {
		return nil, nil
	}

	return respData.Users[0], nil
}
