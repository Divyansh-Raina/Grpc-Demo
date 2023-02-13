package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doGetCredits(c pb.BillingServiceClient) {
	log.Println("doGetCredits was invoked!")
	var uid uint64
	fmt.Print("Enter User ID:- ")
	fmt.Scan(&uid)
	res, err := c.GetCredits(context.Background(), &pb.CreditRequest{Accid: &pb.Account{Id: uid}})
	if err != nil {
		log.Printf("Unable to get credits")
		return
	}
	fmt.Println("----------------")
	fmt.Printf("Id:- %v\nNo of credits:- %v\nExpiry time:- %s\n",
		res.Accid.Id, res.NoOfCredits, res.Expiry.AsTime().String())
	fmt.Println("----------------")

}
