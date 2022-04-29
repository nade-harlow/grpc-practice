package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-practice/gen/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	//defer conn.Close()
	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "hello"})
	if err != nil {
		return
	}
	log.Println(resp)
	log.Println(resp.Msg)
}
