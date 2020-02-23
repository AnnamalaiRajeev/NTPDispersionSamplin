	
	/*
# Authors : [rajeev.chinnikanu@colorado.edu, srsh9018@colorado.edu ]
# Name   : Programming Assignment 1.
# Purpose: gRPC Client
# Date   : Feb2020
# Version : Go 1.13
*/

package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
        "fmt"
        "strings"
	"strconv"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "10.168.0.3:50051"
	defaultName = "world"
)


func create_file(f string) *os.File{
	file, err := os.Create(f)
	if err!=nil {
			panic(err)
	}
	return file
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        fmt.Println("yo")
	if err != nil {
			log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
			name = os.Args[1]
	}
	_ = name


	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        ctx := context.Background()
       	//defer cancel()

	arguments := os.Args
    fmt.Println("hello")	 
    if len(arguments) < 1 {
    	return
    }
    fmt.Println(arguments[1])
	var File_1 string = arguments[1]
	file_object := create_file(File_1)
	file_object.Close()
	
	type clocks_values struct {
		Ti3 int64
		Ti2 int64
		Ti1 int64
		Ti int64
	}
	clock_map := make(map[string]clocks_values)
	

	start := time.Now()
	for {
	var file, err = os.OpenFile(File_1, os.O_RDWR, 0644)
	Ti3 := strconv.FormatInt(time.Now().UnixNano(), 10)
        fmt.Println(Ti3)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: Ti3})
	if err != nil {
			log.Fatalf("could not greet: %v", err)
	}
	s := strings.Split(r.GetMessage(), "???")
	Ti_3, err := strconv.ParseInt(s[0],10,64)
	Ti2, err := strconv.ParseInt(s[1],10,64)
	Ti1, err := strconv.ParseInt(s[2],10,64)
	Ti := time.Now().UnixNano()
        key := time.Now().String()
	clock_map[key] = clocks_values{Ti_3, Ti2, Ti1, Ti}
	jsonString, _ := json.MarshalIndent(clock_map," ","             ")
	file.Write(jsonString)
	err = file.Sync()
	log.Printf("Greeting: %s", r.GetMessage())
	elapsed := time.Since(start)
    if elapsed.Minutes() >= 120{
        err = file.Sync()
        file.Close()
        return
	}
	file.Close()
	time.Sleep(1 * time.Minute)
	}
}
