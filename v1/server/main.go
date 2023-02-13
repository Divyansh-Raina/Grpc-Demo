package main

import (
	pb "draina/demo/v1/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const addr string = "localhost:5002"

type Server struct {
	pb.AboutServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil{
		log.Fatalf("Unable to listen on address: %v\n",err)
	}

	log.Printf("Listening on port:%s",addr)

	//opts := []grpc.ServerOption{};
	tls := true

	if tls{
		
	}
	s := grpc.NewServer()
	pb.RegisterAboutServiceServer(s,Server{})

	// err =  s.Serve(lis) ; err != nil{
	// 	log.Fatalf("%v\n",err)
	// }
}
