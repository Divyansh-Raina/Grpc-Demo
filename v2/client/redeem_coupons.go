package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doRedeemCoupons(c pb.BillingServiceClient) {
	fmt.Println("doRedeemedCoupons was invoked")
	var uid uint64
	var code string
	fmt.Print("Enter User ID:- ")
	fmt.Scan(&uid)
	fmt.Print("Enter Code of the coupon to be used:-")
	fmt.Scan(&code)
	res, err := c.RedeemCoupons(context.Background(), &pb.RedeemRequest{Accid: &pb.Account{Id: uid},Code:code})
	if err != nil {
		log.Print("Was unable to retrieve any used coupons")
		return
	}
	fmt.Print(res.Result)
}
