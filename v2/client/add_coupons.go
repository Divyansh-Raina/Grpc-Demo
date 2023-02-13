package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func doAddCoupons(c pb.BillingServiceClient) {
	log.Printf("doAddCoupons function has been invoked!")

	var uid uint64
	var code string
	var expiry time.Time = time.Now()
	var days int
	var redeem_time time.Time
	var redeemed bool = false
	fmt.Print("Enter User ID:- ")
	fmt.Scan(&uid)
	fmt.Print("Enter Code:- ")
	fmt.Scan(&code)
	fmt.Print("How many days to expire:- ")
	fmt.Scan(&days)
	expiry = expiry.AddDate(0, 0, days)
	redeem_time = expiry

	res := &pb.CouponResponse{
		Accid:      &pb.Account{Id: uid},
		Code:       code,
		Expiry:     timestamppb.New(expiry),
		RedeemTime: timestamppb.New(redeem_time),
		Redeemed:   redeemed,
	}

	result, err := c.AddCoupons(context.Background(), res)
	if err != nil {
		log.Println("Unable to add coupons")
		return
	}
	fmt.Println(result.Result)
}
