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
		fmt.Println("--------- User ------------")
		fmt.Println("Enter (1) to Get User")
		fmt.Println("Enter (2) to Add User")
		fmt.Println("Enter (3) to Update User")
		fmt.Println("Enter (4) to Delete User")
		fmt.Println("--------- Credits ---------")
		fmt.Println("Enter (5) to Get the amount of credits")
		fmt.Println("Enter (6) to Update your credits")
		fmt.Println("-------- Coupons ----------")
		fmt.Println("Enter (7) to Get coupons")
		fmt.Println("Enter (8) to Add coupons")
		fmt.Println("Enter (9) to Redeem your coupons")
		fmt.Println("------- Export ------------")
		fmt.Println("Enter (10) to Export your data to file")
		fmt.Println("--------------------------")
		fmt.Println("Select (0) to Exit")
		fmt.Print("Please enter your choice:- ")
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
		case 10:
			doExportData(c)
		case 0:
			log.Fatalf("Exiting the module!")
		default:
			fmt.Println("Wrong Choice selected, Restarting...")
			continue
		}
	}
}
