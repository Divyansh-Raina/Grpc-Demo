package main

import (
	pb "draina/demo/v2/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//"google.golang.org/grpc/credentials/insecure"
)

const addr string = "localhost:5001"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls{
		certFile := "v2/ssl/ca.crt"
		creds,err := credentials.NewClientTLSFromFile(certFile,"")

		if err != nil{
			log.Fatalf("Error while loading CA trust certificate: %v\n",err)
		}
		opts = append(opts,grpc.WithTransportCredentials(creds))
	}
	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Unable to connect to the server:%v\n", err)
	}
	log.Printf("Connected on Address:%s\n", addr)

	defer conn.Close()

	c := pb.NewBillingServiceClient(conn)

	ServiceHost(c)

}
