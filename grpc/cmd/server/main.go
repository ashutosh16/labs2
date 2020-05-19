// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	"github.com/gopherland/labs2/grpc/internal/generated"
	"github.com/gopherland/labs2/grpc/internal/server"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = "localhost:50052"

//go:generate protoc -I ../../protos --go_out=plugins=grpc:../.. ../../protos/grep.proto
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(server.Measure, server.Logger))
	generated.RegisterGrepperServer(s, server.NewGrepper("assets"))
	reflection.Register(s)

	log.Printf("[Grepper] Listening on %q...", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Boom! Server Failed: %v", err)
	}
}