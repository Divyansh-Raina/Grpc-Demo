package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
	md "draina/demo/v2/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) DeleteUser(ctx context.Context, in *pb.UserRequest) (*emptypb.Empty, error) {
	log.Println("Delete User has been invoked")

	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Unable to Open Databse:%v\n", err)
	}

	var user md.User
	res := db.Where("uid = ?", in.Accid.Id).Delete(&user)
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown error: %v\n", err),
		)
	}
	if res.RowsAffected == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"User was not found",
		)
	}
	return &emptypb.Empty{}, nil
}
