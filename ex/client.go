package ex

import (
	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewClient(addr string) api.DataServiceClient {
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("unable to connect :: %v", err)
	}
	return api.NewDataServiceClient(conn)
}
