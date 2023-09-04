FROM golang:alpine as builder

RUN mkdir -p /tmp/build
WORKDIR /tmp/build

ADD . .

RUN go install github.com/swaggo/swag/cmd/swag
RUN go install github.com/pquerna/ffjson
RUN go mod tidy
RUN go generate ./...

RUN go get -u ./...
RUN go build -o bin/example-api

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /tmp/build/bin/example-api .

EXPOSE 8080
ENTRYPOINT [ "./example-api" ]