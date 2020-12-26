FROM golang AS builder
ADD src/ .
EXPOSE 5000
RUN go get -u github.com/gorilla/mux
RUN GOOS=linux go build -o /app .
CMD ["go", "run", "server.go"]

FROM mongo
COPY --from=builder /app .
CMD ["/app"]

