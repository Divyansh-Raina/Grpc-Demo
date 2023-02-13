package main

import (
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func ServiceHost(c pb.BillingServiceClient) {
	choice := 0
	for {
		fmt.Println("Welcome to Billing Admin Console! Please make your choice")
		fmt.Println("Select 1 for Get User")
		fmt.Println("Select 2 for Add User")
		fmt.Println("Select 3 for Update User")
		fmt.Println("Select 4 for Delete User")
		fmt.Println("Select 5 to Get the amount of credits")
		fmt.Println("Select 6 to Update your credits")
		fmt.Println("Select 7 to Get coupons")
		fmt.Println("Select 8 to Add coupons")
		fmt.Println("Select 9 to Redeem your coupons")
		fmt.Println("Select 0 to Exit")
		fmt.Print("Please select your choice:- ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			doGetUserById(c)
		case 2:
			doAddUser(c)
		case 3:
			doUpdateUser(c)
		case 4:
			doDeleteUser(c)
		case 5:
			doGetCredits(c)
		case 6:
			doUpdateCredits(c)
		case 7:
			doGetCoupons(c)
		case 8:
			doAddCoupons(c)
		case 9:
			doRedeemCoupons(c)
		case 0:
			log.Fatalf("Exiting the module!")
		default:
			fmt.Println("Wrong Choice selected, Restarting...")
			continue
		}
	}
}
