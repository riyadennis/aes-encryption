package main

import (
	"github.com/riyadennis/aes-encryption/ex/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:5052",
		grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cl := api.NewDataServiceClient(conn)
	StoreTest(cl)
}
