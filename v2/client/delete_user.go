package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)
func doDeleteUser(c pb.BillingServiceClient){
	log.Println("doDeleteBlog was invoked!")
	var uid uint64
	fmt.Print("Enter the user id you want to delete:- ")
	fmt.Scan(&uid)
	req := &pb.UserRequest{
		Accid: &pb.Account{Id: uid},
	}
	_,err := c.DeleteUser(context.Background(),req)
	if err != nil{
		log.Fatalf("Error while deleting the value:%v\n",err)
	}
}