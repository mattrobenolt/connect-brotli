package brotli_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/bufbuild/connect-go"
	brotli "go.withmatt.com/connect-brotli"
	pingv1 "go.withmatt.com/connect-brotli/internal/gen/connect/ping/v1"
	"go.withmatt.com/connect-brotli/internal/gen/connect/ping/v1/pingv1connect"
)

func ExampleNew() {
	// Get client and server options
	clientOpts, serverOpts := brotli.New()

	// Create a server.
	_, h := pingv1connect.NewPingServiceHandler(&pingServer{}, serverOpts)
	srv := httptest.NewServer(h)
	client := pingv1connect.NewPingServiceClient(
		http.DefaultClient,
		srv.URL,
		clientOpts,
		// Compress requests with Brotli.
		connect.WithSendCompression(brotli.Name),
	)
	req := connect.NewRequest(&pingv1.PingRequest{
		Number: 42,
	})
	req.Header().Set("Some-Header", "hello from connect")
	res, err := client.Ping(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("The answer is", res.Msg)
	fmt.Println(res.Header().Get("Some-Other-Header"))
	//OUTPUT:
	//hello from connect
	//The answer is number:42
	//hello!
}

type pingServer struct {
	pingv1connect.UnimplementedPingServiceHandler // returns errors from all methods
}

func (ps *pingServer) Ping(
	ctx context.Context,
	req *connect.Request[pingv1.PingRequest],
) (*connect.Response[pingv1.PingResponse], error) {
	// connect.Request and connect.Response give you direct access to headers and
	// trailers. No context-based nonsense!
	fmt.Println(req.Header().Get("Some-Header"))
	res := connect.NewResponse(&pingv1.PingResponse{
		// req.Msg is a strongly-typed *pingv1.PingRequest, so we can access its
		// fields without type assertions.
		Number: req.Msg.Number,
	})
	res.Header().Set("Some-Other-Header", "hello!")
	return res, nil
}
