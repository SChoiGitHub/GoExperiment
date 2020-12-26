FROM golang AS builder
RUN go get -u github.com/gorilla/mux
ADD src/ .
ADD internal/ /go/src/repository/
EXPOSE 5000
RUN GOOS=linux go build -mod=vendor -o /app .
CMD ["go", "run", "server.go"]

FROM mongo
COPY --from=builder /app .
CMD ["/app"]

