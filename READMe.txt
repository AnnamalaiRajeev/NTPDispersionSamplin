# Authors
	/*
# Authors : [rajeev.chinnikanu@colorado.edu, srsh9018@colorado.edu ]
# Name   : Programming Assignment 1.
# Purpose: UDPServer
# Date   : Feb2020
# Version : Go 1.13
*/



## Configuration

[1] Download Go build from https://golang.org/dl/ Add Go project path to you PATH Variable
[2] Test go using Prompt# go version.
[2] Make a folder named "src" in your project directory.
[3] Unzip the file into the Src directory.
[4] Install gRPC on your terminal using 
--> go get -u google.golang.org/grpc
--> go get -u github.com/golang/protobuf/protoc-gen-go
--> export PATH=$PATH:$GOPATH/bin
[5] You are set to run the Make File.



### Run

[1] Servers listen on 0.0.0.0:**** where port number (****) will vary on the version of server running [gRPC vs UDP]

--> gRPC server listens on 0.0.0.0:50051
--> UDP server Listens on 0.0.0.0:3333

To Run gRPC setup:
[1] cd to client Directory and run  --> make grpc_server
[2] cd to server directory and run  --> make grpc_client 1="x.x.x.x:yyyy" 2="filename.json"

To Run UDP setup:
[1] cd to client Directory and run  --> make udp_server
[2] cd to server directory and run  --> make udp_client 1="x.x.x.x:yyyy" 2="filename.json"





