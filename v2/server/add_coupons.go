package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	md "draina/demo/v2/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) AddCoupons(ctx context.Context, in *pb.CouponResponse) (*pb.OkMessage, error) {
	log.Printf("Add Coupons function was invoked!:%v\n", in)
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not open the database:%v\n", err)
	}
	db.AutoMigrate(md.Coupon{})
	coupon := md.Coupon{
		Uid:         in.Accid.Id,
		Code:        in.Code,
		Expiry:      in.Expiry.AsTime(),
		Redeem_Time: in.RedeemTime.AsTime(),
		Redeemed:    in.Redeemed,
	}
	var c2 md.Coupon
	res := db.Where("uid = ? and code = ?", in.Accid.Id, in.Code).Find(&c2)
	if res.RowsAffected != 0 {
		return &pb.OkMessage{
			Result: "Coupon already Exists",
		},nil
	}
	res = db.Create(&coupon)
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Error:%v\n", err),
		)
	}
	return &pb.OkMessage{
		Result: "Coupon has been added",
	}, nil
}
