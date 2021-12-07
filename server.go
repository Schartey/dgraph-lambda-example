package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgraph-io/dgo/v210"
	dapi "github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi"
	"github.com/machinebox/graphql"
	"google.golang.org/grpc"

	"github.com/schartey/dgraph-lambda-example/lambda/generated"
	"github.com/schartey/dgraph-lambda-example/lambda/resolvers"
	"github.com/schartey/dgraph-lambda-go/api"
)

func main() {

	r := chi.NewRouter()

	dgraphUrl := os.Getenv("DGRAPH_URL")
	fmt.Println(dgraphUrl)

	resolver := setupResolver(dgraphUrl)
	executer := generated.NewExecuter(resolver)
	lambda := api.New(executer)

	r.Get("/health", func(rw http.ResponseWriter, r *http.Request) {
		type health struct {
			Success bool   `json:"success"`
			Message string `json:"message"`
		}
		h := health{Success: true, Message: "Lambda is working"}
		res, err := json.Marshal(&h)
		if err != nil {
			rw.WriteHeader(500)
		}
		rw.Write(res)
	})
	r.Post("/graphql-worker", lambda.Route)

	fmt.Println("Lambda listening on 8686")
	fmt.Println(http.ListenAndServe(":8686", r))

}

func setupResolver(dgraphUrl string) *resolvers.Resolver {
	gql := createGraphQLClient(dgraphUrl)
	dql := createDGraphClient(dgraphUrl)

	return &resolvers.Resolver{
		Gql: gql,
		Dql: dql,
	}
}

func createGraphQLClient(dgraphUrl string) *graphql.Client {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	client := graphql.NewClient(fmt.Sprintf("http://%s:8080/graphql", dgraphUrl))
	client.Log = func(s string) { log.Println(s) }
	return client
}

func createDGraphClient(dgraphUrl string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(fmt.Sprintf("%s:9080", dgraphUrl), grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
	}

	return dgo.NewDgraphClient(
		dapi.NewDgraphClient(d),
	)
}
