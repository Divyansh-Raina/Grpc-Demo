package main

import (
	"context"
	md "draina/demo/v2/model"
	pb "draina/demo/v2/proto"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) GetCredits(ctx context.Context, in *pb.CreditRequest) (*pb.CreditResponse, error) {
	log.Printf("GetCredits function was invoked!:%v\n", in)

	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to retrieve database")
	}
	if !checkUser(in.Accid.Id) {
		return nil, status.Errorf(
			codes.NotFound,
			"User not found",
		)
	}
	db.AutoMigrate(md.Credit{})
	var credits md.Credit
	res := db.Where("uid = ?", in.Accid.Id).First(&credits)
	if res.RowsAffected == 0 {
		credits = md.Credit{Uid: in.Accid.Id, No_of_credits: 0, Expiry: time.Now()}
		db.Create(&credits)
	}
	return &pb.CreditResponse{
		Accid:       &pb.Account{Id: credits.Uid},
		NoOfCredits: credits.No_of_credits,
		Expiry:      timestamppb.New(credits.Expiry),
	}, nil
}
