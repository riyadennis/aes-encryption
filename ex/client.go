package ex

import (
	"os"

	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewClient() api.DataServiceClient {
	var conn *grpc.ClientConn
	var err error
	if os.Getenv("PORT") == "" {
		logrus.Fatal("no port found")
	}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err = grpc.Dial("127.0.0.1:5300", opts...)
	if err != nil {
		logrus.Fatalf("unable to connect :: %v", err)
	}
	return api.NewDataServiceClient(conn)
}
