package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doGetUserById(c pb.BillingServiceClient) {
	log.Println("GetUserByID is invoked")

	var Id uint64

	fmt.Print("Enter ID to be retrieved:- ")
	fmt.Scan(&Id)
	//fmt.Scanf(" Id:- %d\n", &Id)
	accid := &pb.Account{
		Id: Id,
	}
	req := &pb.UserRequest{
		Accid: accid,
	}
	res, err := c.GetUserById(context.Background(), req)
	if err != nil {
		log.Fatalf("could not retrieve details from the server")
	}
	fmt.Println("----------------")
	fmt.Printf(" Id:- %d\n Name:-%s\n Age:- %d\n Type of Account:- %s\n", res.Accid.Id, res.Name, res.Age, res.TypeOfAcc)
	fmt.Println("----------------")

}
