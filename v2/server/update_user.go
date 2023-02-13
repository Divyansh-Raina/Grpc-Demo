package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"

	//"google.golang.org/grpc/internal/status"
	md "draina/demo/v2/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) UpdateUser(ctx context.Context, in *pb.UserResponse) (*emptypb.Empty, error) {
	log.Println("Update User was invoked")
	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database:%v\n", err)
	}
	// if !checkUser(in.Accid.Id) {
	// 	return nil, status.Errorf(
	// 		codes.NotFound,
	// 		"User not found",
	// 	)
	// }
	db.AutoMigrate(md.User{})
	var user md.User
	res := db.Model(&user).Where("uid = ?", in.Accid.Id).Updates(md.User{Name: in.Name, Age: in.Age, Type_of_acc: in.TypeOfAcc})
	if res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Could not find the user",
		)
	}
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown error:%v\n", res.Error),
		)
	}

	return &emptypb.Empty{}, nil
}
