package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	 md "draina/demo/v2/model"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) GetUserById(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Println("GetUser Service is invoked:- ")

	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	db.AutoMigrate(&md.User{})
	var user md.User
	res := db.First(&user, "uid = ?", in.Accid.Id)

	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown Error:%v", err),
		)
	}
	if res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"User not found",
		)
	}
	return &pb.UserResponse{
		Accid:     &pb.Account{Id: user.Uid},
		Name:      user.Name,
		Age:       user.Age,
		TypeOfAcc: user.Type_of_acc,
	}, nil
}
