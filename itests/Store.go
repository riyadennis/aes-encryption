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
			Message: `This guide describes how CircleCI 
finds and runs config.yml and how you can use shell 
commands to do things, then it outlines how config.yml 
can interact with code and kick-off a build followed by
how to use docker containers to run in precisely the 
environment that you need. Finally, there is a short
exploration of work flows so you can learn to orchestrate your build,
tests, security scans, approval steps, and deployment.
CircleCI believes in configuration as code. As a result, 
the entire delivery process from build to deploy is orchestrated
through a single file called config.yml. 
The config.yml file is located in a folder called .circleci at 
the top of your project. CircleCI uses the YAML syntax for config,
see the Writing YAML document for basics.
`,
		},
	}
	_, err := cl.Store(ctx, in)
	if err != nil {
		logrus.Fatal(err)
	}
}
