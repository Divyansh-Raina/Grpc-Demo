package main

import (
	"context"
	md "draina/demo/v2/model"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) UpdateCredits(ctx context.Context, in *pb.CreditResponse) (*pb.OkMessage, error) {
	log.Println("Update Credits function is invoked!")
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
	db.AutoMigrate(md.Credit{})
	res := db.Model(md.Credit{}).Where("uid = ?", in.Accid.Id).Updates(md.Credit{No_of_credits: in.NoOfCredits, Expiry: in.Expiry.AsTime()})

	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Error:%v\n", res.Error),
		)
	}
	if res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"No User found having credits",
		)
	}

	return &pb.OkMessage{
		Result: "Update was Successful!",
	}, nil
}
