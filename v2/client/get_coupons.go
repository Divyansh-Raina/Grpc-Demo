package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"io"
	"log"
)

func doGetCoupons(c pb.BillingServiceClient) {
	fmt.Println("DoGetCoupons function was invoked!")

	var uid uint64
	fmt.Print("Enter the User ID:- ")
	fmt.Scan(&uid)
	req := &pb.CouponRequest{
		Accid: &pb.Account{Id: uid},
	}
	stream, err := c.GetCoupons(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the function%v\n", err)

	}
	var count int32 = 0
	for {

		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("-----------------")
			fmt.Println("Total No. of coupons:- ", count)
			fmt.Println("-----------------")
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the data stream:- %v\n", err)
		}
		fmt.Println("-----------------")
		fmt.Println("Code:- ", msg.Code)
		fmt.Println("Expiry:- ", msg.Expiry.AsTime().String())
		fmt.Println("Redeem Time:- ", msg.RedeemTime.AsTime().String())
		fmt.Println("Redeeemed:- ", msg.Redeemed)
		count++
	}

}
