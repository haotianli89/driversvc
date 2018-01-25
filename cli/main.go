package main

import (
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/haotianli89/driversvc/pb"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	client := driversvc.NewDriversvcClient("driversvc", client.DefaultClient)

	rsp, err := client.GetDrivers(context.Background(), &driversvc.GetDriversRequest{/*Id: "add45514-7dfb-4ce7-a74b-26d380939833"*/})
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(rsp)
	}
}