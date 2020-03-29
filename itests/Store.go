package main

import (
	"context"

	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
)

func StoreTest(cl api.DataServiceClient) {
	ctx := context.Background()
	in := &api.DataRequest{
		Data: &api.Data{
			Message: "Test message",
		},
	}
	_, err := cl.Store(ctx, in)
	if err != nil {
		logrus.Fatal(err)
	}
}
