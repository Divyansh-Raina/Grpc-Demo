package main

import (
	"context"
	md "draina/demo/v2/model"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"

	//"os"

	//"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) RedeemCoupons(ctx context.Context, in *pb.RedeemRequest) (*pb.OkMessage, error) {
	log.Println("Redeemed Coupons is Invoked!")
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not open the database:%v\n", err)
	}
	if !checkUser(in.Accid.Id) {
		return nil, status.Errorf(
			codes.NotFound,
			"User not found",
		)
	}
	db.AutoMigrate(md.Coupon{})

	var coupons []md.Coupon
	db.Where("uid = ? and code = ?", in.Accid.Id,in.Code).Find(&coupons)
	res := db.Where("uid = ? and code = ?", in.Accid.Id, in.Code).Delete(&md.Coupon{})
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Error:%v\n", err),
		)
	}
	// fmt.Print("Do you want to save redeemed coupons in a file? (y/N)")
	// var flag string
	// fmt.Scan(&flag)
	// if flag == "y" {
	// 	err = CreateCouponFile(in.Accid.Id,coupons)
	// 	if err != nil {
	// 		log.Printf("Wasnt able to write a file with that ID\nAborting...")
	// 	}
	// }
	return &pb.OkMessage{
		Result: fmt.Sprintf("%v coupons have been redeemed", res.RowsAffected),
	}, nil
}

// func CreateCouponFile(id uint64, cps []md.Coupon) error {
// 	file, err := os.Create(fmt.Sprintf("%v.txt", id))
// 	if err != nil {
// 		return status.Errorf(
// 			codes.AlreadyExists,
// 			"File already exists",
// 		)
// 	}
// 	for _, c := range cps {
// 		file.WriteString(fmt.Sprintf("Id:%v\nCode:%s\nExpiry:%v\nRedeem Time:%v\nRedeemed:%t\n------\n",
// 			id, c.Code, c.Expiry, c.Redeem_Time, c.Redeemed))
// 	}
// 	return nil
// }
