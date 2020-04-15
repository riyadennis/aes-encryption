package main

import (
	"context"

	"github.com/riyadennis/aes-encryption/ex/api"
	"github.com/sirupsen/logrus"
)

var message = `This guide describes how CircleCI 
finds and runs config.yml and how you can use shell 
commands to do things, then it outlines how config.yml 
can interact with code and kick-off a build followed by
how to use docker containers to run in precisely the 
environment that you need. Finally, there is a short
exploration of work flows so you can learn to orchestrate your build,
tests, security scans, approval steps, and deployment.
`

func StoreRetrieveTest(cl api.DataServiceClient) {
	ctx := context.Background()
	in := &api.DataRequest{
		Data: &api.Data{
			Message: message,
		},
	}
	resp, err := cl.Store(ctx, in)
	if err != nil {
		logrus.Fatalf("unable to store :: %v", err)
	}
	req := &api.RetrieveRequest{
		EncryptionId:  resp.EncryptionId,
		EncryptionKey: resp.EncryptionKey,
	}
	rResp, err := cl.Retrieve(ctx, req)
	if err != nil {
		logrus.Fatal(err)
	}
	if rResp.Data.Message != message {
		logrus.Fatal("invalid message")
	}
}
