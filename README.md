# connect-brotli

[![GoDoc](https://pkg.go.dev/badge/go.withmatt.com/connect-brotli.svg)](https://pkg.go.dev/go.withmatt.com/connect-brotli)

This package provides improved compression schemes for [Buf Connect](https://github.com/bufbuild/connect-go).

Compression is provided from the [github.com/andybalholm/brotli](https://github.com/andybalholm/brotli) package.

# Usage

The `brotli.New` function will return an option that allow both client and servers to compress and decompress using brotli.

```go
import "go.withmatt.com/connect-brotli"

opts := brotli.New()
```

To enable client compression and force a specific method use `connect.WithSendCompression(brotli.Name)`.

For more details and options see the [documentation](https://pkg.go.dev/go.withmatt.com/connect-brotli).
