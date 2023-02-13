package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doUpdateUser(c pb.BillingServiceClient) {
	log.Println("doUpdateUser was invoked!")
	var uid uint64
	fmt.Print("Enter User Id to update:-")
	fmt.Scan(&uid)
	Accid := pb.UserRequest{
		Accid: &pb.Account{
			Id: uid,
		},
	}
	res, err := c.GetUserById(context.Background(), &Accid)
	if err != nil {
		log.Fatalf("Unable to access User data:%v\n", err)
	}
	var name string
	var age uint32
	var type_of_acc string
	fmt.Print("Name:- ")
	fmt.Scan(&name)
	fmt.Print("Age:- ")
	fmt.Scan(&age)
	fmt.Print("Type of Account:- ")
	fmt.Scan(&type_of_acc)
	if len(name) == 0 {
		name = res.Name
	}
	if age < 18 {
		if res.Age > 18 {
			age = res.Age
		} else {
			age = 18
		}
	}
	//type of account
	_, err = c.UpdateUser(context.Background(), &pb.UserResponse{
		Accid:     &pb.Account{Id: uid},
		Age:       age,
		Name:      name,
		TypeOfAcc: type_of_acc,
	})
	if err != nil {
		log.Printf("Error while updating:%v\n", err)
		return
	}
	fmt.Println("Updates were successful")
}
