package main

import (
	"context"
	"fmt"

	"github.com/riyadennis/aes-encryption/ex/api"
	client2 "github.com/riyadennis/aes-encryption/ex/client"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := api.NewDataServiceClient(conn)
	client := client2.AesClient{}
	req := client.DataRequest("test pay load", "test encryption")
	resp, err := c.Store(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp)
}
