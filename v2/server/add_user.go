package main

import (
	"context"
	md "draina/demo/v2/model"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"

	//"sync"
	//"google.golang.org/grpc/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *Server) AddUser(ctx context.Context, in *pb.UserResponse) (*pb.OkMessage, error) {

	db, err := gorm.Open(sqlite.Open("demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to retrieve the database: %v\n", err)
	}

	db.AutoMigrate(&md.User{})
	//var m sync.Mutex
	//go func() {
	//m.Lock()
	var user md.User
	
	user = md.User{
		Uid:         in.Accid.Id,
		Name:        in.Name,
		Age:         in.Age,
		Type_of_acc: in.TypeOfAcc}
	res := db.Where("uid = ?", user.Uid).First(&md.User{})
	if res.RowsAffected != 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			"User ID already exists",
		)
	}
	user = md.User{
		Uid:         in.Accid.Id,
		Name:        in.Name,
		Age:         in.Age,
		Type_of_acc: in.TypeOfAcc}

	res = db.Create(&user)
	//m.Unlock()
	if res.Error != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown error:%v", res.Error),
		)
	}
	//}()
	// db.Where("uid = ?", user.Uid).First(&user)
	// log.Println(user.Uid, user.Name)
	//var idx uint64 = *&in.Accid.Id

	return &pb.OkMessage{
		Result: fmt.Sprintf("User was created with ID %v\n", user.Uid),
	}, nil
}
