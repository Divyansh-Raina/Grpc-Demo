package main

import (
	"context"
	pb "draina/demo/v2/proto"
	"fmt"
	"log"
)

func doExportData(c pb.BillingServiceClient) {
	log.Println("Export data function was invoked!")
	var uid uint64
	fmt.Print("Enter the user ID of the person to export data:- ")
	fmt.Scan(&uid)
	res, err := c.ExportData(context.Background(), &pb.Account{Id: uid})
	if err != nil {
		log.Printf("Unsuccessful export")
		return
	}
	fmt.Println(res.Result)

}
