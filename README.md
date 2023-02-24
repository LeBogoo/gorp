# Gorp

Gorp is a very basic HTTP reverse proxy. It is meant to be used as a simple way to proxy requests to a backend server. It is not meant to be a full featured proxy.

> DISCLAIMER: This project is not meant to be used in production. It is meant to be for educational purposes only!

## Creating a proxy

Create a config file in the config directory. The name of the file is irrelevant. The config file should be a JSON file with the following format:

```json
{
    "host": "subdomain.example.com",
    "server": "http://hostname:8080"
}
```

The `host` field is the url from which the proxy will be accessed. The `server` field is the url of the server to which the proxy will forward requests. Note that the `server` field must be a valid HTTP url.

The proxy doesn't have to be restarted when a new config file is added.

If a request is made to the proxy it will scan the config directory for a config file where the `host` field matches the request's host. If a config file is found, the request will be forwarded to the specified server.

## Running the proxy using go

To run the proxy using go, run the following command:

```bash
go run main.go
```

## Building and running the proxy

To build the proxy, run the following command:

```bash
go build
```

This will create a binary called `gorp` or on Windows `gorp.exe`. To run the proxy, run the following command:

```bash
./gorp
```
