FROM golang
ADD src/ .
EXPOSE 5000
RUN go get -u github.com/gorilla/mux
CMD ["go", "run", "server.go"]

