package main

import (
	"context"
	md "draina/demo/v2/model"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) ExportData(ctx context.Context, in *pb.Account) (*pb.OkMessage, error) {
	log.Println("Export data has been invoked with", in.Id)
	//filename := fmt.Sprintf("%v.txt", in.Id)
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	//User Export
	db.AutoMigrate(&md.User{})
	var user md.User
	res := db.First(&user, "uid = ?", in.Id)
	CreateUserFile(in.Id, user)
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to find User:%v\n", res.Error),
		)
	}
	// Credit Export
	db.AutoMigrate(&md.Credit{})
	var credit md.Credit
	res = db.First(&user, "uid = ?", in.Id)
	CreateCreditsFile(in.Id, credit)
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to find credits for given user id:%v\n", res.Error),
		)
	}
	// Coupon Export
	db.AutoMigrate(&md.Coupon{})
	var coupons []md.Coupon
	res = db.Where("uid = ?", in.Id).Find(&coupons)
	CreateCouponFile(in.Id, coupons)
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable to find coupons for given user id:%v\n", res.Error),
		)
	}

	return &pb.OkMessage{
		Result: fmt.Sprintf("Export Successful with User ID:- %v\n", in.Id),
	}, nil
}

func CreateCouponFile(id uint64, cps []md.Coupon) error {
	//pwd, _ := os.Getwd()
	//path := filepath.Join(pwd, "data", fmt.Sprint(fmt.Sprint(id)+"_Coupons.txt"))
	file, err := os.Create(fmt.Sprint(id) + "_Coupons.txt")

	//stats,err := os.Stat(fn)
	if err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			"File already exists",
		)
	}
	defer file.Close()
	for _, c := range cps {
		file.WriteString(fmt.Sprintf("Id:%v\nCode:%s\nExpiry:%v\nRedeem Time:%v\nRedeemed:%t\n------\n",
			id, c.Code, c.Expiry.String(), c.Redeem_Time.String(), c.Redeemed))
	}
	return nil
}
func CreateCreditsFile(id uint64, c md.Credit) error {
	//pwd, _ := os.Getwd()
	//path := filepath.Join(pwd, "data", fmt.Sprint(fmt.Sprint(id)+"_Credits.txt"))
	file, err := os.Create(fmt.Sprint(id) + "_Credits.txt")
	if err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			"File already exists",
		)
	}
	defer file.Close()
	//stats,err := os.Stat(fn)
	if err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			"File already exists",
		)
	}

	file.WriteString(fmt.Sprintf("Id:%v\nNo of Credits:%v\nExpiry:%v\n",
		id, c.No_of_credits, c.Expiry.String()))

	return nil
}
func CreateUserFile(id uint64, c md.User) error {
	//pwd, _ := os.Getwd()
	//path := filepath.Join(pwd, "data", fmt.Sprint(fmt.Sprint(id)+"_User.txt"))
	file, err := os.Create(fmt.Sprint(id) + "_User.txt")
	if err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			"File already exists",
		)
	}
	defer file.Close()
	//stats,err := os.Stat(fn)
	if err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			"File already exists",
		)
	}
	file.WriteString(fmt.Sprintf("Id:%v\nName:%s\nAge:%v\nType of Account:%v\n",
		id, c.Name, c.Age, c.Type_of_acc))

	return nil
}
