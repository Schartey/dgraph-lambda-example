# dgraph-lambda-example

Example repository for dgraph-lambda-go.

## Generate resolvers

go run github.com/schartey/dgraph-lambda-go generate

## DGraph

Update schema on dgraph

```
curl -X POST localhost:8080/admin/schema --data-binary '@schema.graphql'
```

## Build/Push image

docker build -t your/tag .

docker push your/tag