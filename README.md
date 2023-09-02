# example-api

## Build

### Local

```bash
go get -u ./...
go install github.com/swaggo/swag/cmd/swag
go install github.com/pquerna/ffjson

go generate ./...
go build -o bin/example-api 

```

### Container (Recommended)

```bash
docker build -t example-api .

```

## Run

### Local

```bash
chmod +x ./bin/example-api
./example-api --logLevel=DEBUG
# ./example-api --logLevel=DEBUG --port=5000 --> on custom port
```


### Container (Recommended)

```bash
docker run -p 8080:8080 example-api --logLevel=DEBUG
# docker run -p 5000:5000 example-api --logLevel=DEBUG --port=5000 --> on custom port
```

## Usage

See `http://localhost:8080/docs` --> Swagger Documentation
