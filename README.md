# Forward Proxy Example

This is just simple example of a forward proxy working using [goproxy](https://github.com/elazarl/goproxy/).

It consists of a proxy server, a server and a client all of which are contained
in their own folders.

## Instructions

Start the proxy server
```go
cd proxy

go run main.go
```

Start the server
```go
cd server
go run main.go
```

Run the client
```go
cd client
go run main.go
```

## Expected output

To prove the requests are going via the proxy, it appends on the header:
`"X-GoProxy:hello world"` which should be visible in the servers logs.
