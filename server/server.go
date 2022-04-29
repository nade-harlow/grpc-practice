package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-practice/gen/proto"
	"log"
	"net"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (r *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
