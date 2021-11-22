package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	fibonaccipb "github.com/duckladydinh/fibonacci-go-gateway/gen/go/giathuan/examples/fibonacci/v1"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9999", "gRPC server endpoint")
	httpServerEndpoint = flag.String("http-server-endpoint", "localhost:10000", "HTTP server endpoint")
)

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register gRPC server endpoint.
	log.Printf("Registering gRPC server at %v\n", *grpcServerEndpoint)
	mux := runtime.NewServeMux()
	if err := fibonaccipb.RegisterFibonacciServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint).
	log.Printf("Starting HTTP server at %v\n", *httpServerEndpoint)
	return http.ListenAndServe(*httpServerEndpoint, mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
