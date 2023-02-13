package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func doUpdateCredits(c pb.BillingServiceClient) {
	fmt.Println("doUpdateCredits function was invoked!")
	var uid uint64
	var creds int32
	fmt.Print("Enter Account ID to be updated:-")
	fmt.Scan(&uid)
	fmt.Print("Enter no of credits you want to add or subtract:- ")
	fmt.Scan(&creds)
	res := &pb.CreditResponse{
		Accid: &pb.Account{
			Id: uid,
		},
		NoOfCredits: creds,
		Expiry:      timestamppb.New(time.Now().AddDate(1, 0, 0)),
	}
	_, err := c.UpdateCredits(context.Background(), res)
	if err != nil {
		log.Printf("Could not update the credits")
	}
}
