package main

import (
	"log"
	pb "draina/demo/v2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr string = "localhost:5001"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect to the server:%v\n", err)
	}
	log.Printf("Connected on Address:%s\n", addr)

	defer conn.Close()

	c := pb.NewBillingServiceClient(conn)

	ServiceHost(c)

}
