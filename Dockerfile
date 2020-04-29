FROM golang:1.11.1-alpine
EXPOSE 8080
RUN apk add --update git; \
    mkdir -p ${GOPATH}/go-rest-api-docker; \
    go get -u github.com/gorilla/mux
WORKDIR ${GOPATH}/go-rest-api-docker/
COPY go-rest-api-docker.go ${GOPATH}/go-rest-api-docker/
RUN go build -o go-rest-api-docker .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/go-rest-api-docker/go-rest-api-docker .
CMD [ "/app/go-rest-api-docker" ]
