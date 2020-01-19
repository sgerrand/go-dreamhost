# go-dreamhost

`go-dreamhost` is a Go library for accessing the [Dreamhost
API](https://help.dreamhost.com/hc/en-us/articles/217560167).

## Usage

```go
import "github.com/sgerrand/go-dreamhost"
```

Construct a new Dreamhost client. For example:

```go
client := dreamhost.NewClient("your-api-key", nil)
```

## Authentication

The `go-dreamhost` client will pass through the API key provided as part
of creating a new client.

## License

This library is distributed under the [MIT license].
