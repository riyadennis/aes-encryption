package ex

import (
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
)

func NewClient() api.DataServiceClient {
	var conn *grpc.ClientConn
	var err error
	if os.Getenv("PORT") == "" {
		logrus.Fatal("no port found")
	}
	conn, err = grpc.Dial("localhost"+os.Getenv("PORT"), grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("unable to connect :: %v", err)
	}
	return api.NewDataServiceClient(conn)
}
