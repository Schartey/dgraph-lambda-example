FROM golang:1.16

WORKDIR /lambda

COPY . .

EXPOSE 8686

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#CMD [ "go", "build", "server.go"]
ENTRYPOINT CompileDaemon -log-prefix=true -build="go build -o server" -command="./server"