FROM golang AS builder
RUN go get -u github.com/gorilla/mux
RUN go get -u go.mongodb.org/mongo-driver/mongo
COPY src /go/src
RUN GOOS=linux go build -o /app src/server.go
CMD ["/app"]

