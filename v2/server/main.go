package main

import (
	pb "draina/demo/v2/proto"
	"log"
	"net"
	//"time"
	//"draina/demo/v2/model"
	"google.golang.org/grpc"
	//"gorm.io/gorm"
)

const addr string = "localhost:5001"

type Server struct {
	pb.BillingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to connect : %v\n", err)
	}
	log.Printf("Listening on address: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBillingServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Unable to serve on %v\n", err)
	}

}
