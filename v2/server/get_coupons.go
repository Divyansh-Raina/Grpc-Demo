package main

import (
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	md "draina/demo/v2/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) GetCoupons(in *pb.CouponRequest, stream pb.BillingService_GetCouponsServer) error {
	log.Println("Get Coupons function is invoked!")

	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to fetch Databse:%v\n", err)
	}
	if !checkUser(in.Accid.Id) {
		return status.Errorf(
			codes.NotFound,
			"User not found",
		)
	}
	db.AutoMigrate(&md.Coupon{})
	var coupons []md.Coupon
	res := db.Where("uid = ?", in.Accid.Id).Find(&coupons)
	if res.Error != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to find coupons for user id:%v\n", res.Error),
		)
	}
	for _, coup := range coupons {
		stream.Send(&pb.CouponResponse{
			Accid:      &pb.Account{Id: coup.Uid},
			Code:       coup.Code,
			Expiry:     timestamppb.New(coup.Expiry),
			RedeemTime: timestamppb.New(coup.Redeem_Time),
			Redeemed:   coup.Redeemed,
		})
	}
	return nil
}
