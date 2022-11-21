package brotli

import (
	"github.com/andybalholm/brotli"
	"github.com/bufbuild/connect-go"
)

const (
	BestSpeed          = brotli.BestSpeed
	BestCompression    = brotli.BestCompression
	DefaultCompression = brotli.DefaultCompression
)

const Name = "br"

// New returns client and handler options for the brotli compression method
// using the default compression level.
func New() connect.Option {
	return NewWithLevel(DefaultCompression)
}

// NewWithLevel returns client and handler options for the brotli compression
// method for your prefered compression level.
func NewWithLevel(level int) connect.Option {
	d, c := brComp(level)
	return compressorOption{
		ClientOption:  connect.WithAcceptCompression(Name, d, c),
		HandlerOption: connect.WithCompression(Name, d, c),
	}
}

func brComp(level int) (func() connect.Decompressor, func() connect.Compressor) {
	d := func() connect.Decompressor { return &brrWrapper{brotli.NewReader(nil)} }
	c := func() connect.Compressor { return brotli.NewWriterLevel(nil, level) }
	return d, c
}

type compressorOption struct {
	connect.ClientOption
	connect.HandlerOption
}

type brrWrapper struct{ *brotli.Reader }

func (b *brrWrapper) Close() error { return b.Reset(nil) }
