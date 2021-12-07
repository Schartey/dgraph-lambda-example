# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o /lambda

FROM alpine:3.14

WORKDIR /

COPY --from=build /lambda /lambda

EXPOSE 8686

ENTRYPOINT ["/lambda"]