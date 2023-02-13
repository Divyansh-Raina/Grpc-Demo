package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doAddUser(c pb.BillingServiceClient) {
	fmt.Println("Add User Function was invoked!")
	var uid uint64
	var name string
	var age uint32
	var type_of_acc string
	fmt.Print("Id:- ")
	fmt.Scan(&uid)
	fmt.Print("Name:- ")
	fmt.Scan(&name)
	fmt.Print("Age:- ")
	fmt.Scan(&age)
	fmt.Print("Type of Account:- ")
	fmt.Scan(&type_of_acc)
	//fmt.Scanf(" Id:- %d\n Name:-%s\n Age:- %d\n Type of Account:- %s\n", &Id, &name, &age, &type_of_acc)
	accid := &pb.Account{
		Id: uid,
	}
	if age < 18{
		log.Printf("Invalid age")
	}
	
	res, err := c.AddUser(context.Background(), &pb.UserResponse{
		Accid:     accid,
		Name:      name,
		Age:       age,
		TypeOfAcc: type_of_acc,
	})
	if err != nil {
		log.Fatalf("could not add user to the server")
	}
	fmt.Printf(res.Result)
}
