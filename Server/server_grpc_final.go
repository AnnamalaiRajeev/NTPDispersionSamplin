
	/*
# Authors : [rajeev.chinnikanu@colorado.edu, srsh9018@colorado.edu ]
# Name   : Programming Assignment 1.
# Purpose: gRPCServer
# Date   : Feb2020
# Version : Go 1.13
*/

// Package main implements a server for Greeter service.
package main

import (
        "context"
        "strconv"
        "log"
        "net"
        "time"
        "google.golang.org/grpc"
        pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
        port = "0.0.0.0:50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
        pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        log.Printf("Received: %v", in.GetName())
        Ti3 := in.GetName()
        Ti2 := strconv.FormatInt(time.Now().UnixNano(), 10)
        time.Sleep(1 * time.Second)
        Ti1 := strconv.FormatInt(time.Now().UnixNano(), 10)
        return &pb.HelloReply{Message: Ti3 + "???" + Ti2 + "???" + Ti1}, nil
}

func main() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterGreeterServer(s, &server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
