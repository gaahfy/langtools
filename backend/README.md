# Langtools backend

## How to run the backend service locally

### Build the server

```shell
go build server.go
```

### Run the server

```shell
LANGTOOLS_BACKEND_PORT=":8080" ./server
```

Alternatively, `LANGTOOLS_BACKEND_PORT` can be set as an environment variable of
the host.