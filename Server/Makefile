#makefile#
Help:
	echo "Client takes two arguments Server_IP and File.json. I.e make udp_client 1=8.8.8.8 2=example.json"

grpc_server:
	echo "Running gRPC Server. Server listening on 0.0.0.0:50051"
	go run server_grpc_final.go

grpc_client:
	echo "Running Client"
	go run client_grpc_final.go $(1) $(2)

udp_server:
	echo "Running udp Server .Server will listen to 0.0.0.0:3333"
	go run server.go

udp_client:
	echo "Running udp Client"
	go run client.go $(1) $(2)

