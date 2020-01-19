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

## Configuration

Follow the following steps to generate a new API key:
1. Visit [Dreamhost's Web Panel](https://panel.dreamhost.com/?tree=home.api).
1. Enter a comment for this key. I suggest adding one which will make it easy
   to identify later.
1. Select the functions which this API key should have access to. I suggest
   limiting these to those which you want to control.
1. Click "Generate a new API Key now!"

## Authentication

The `go-dreamhost` client will pass through the API key provided as part
of creating a new client.

## License

This library is distributed under the [MIT license](LICENSE).
