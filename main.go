package main

import (
	"context"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/aes-encryption/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	errChan := make(chan error)
	go server.Run(ctx, errChan)
	for {
		select {
		case <-ctx.Done():
			break
		case err := <-errChan:
			logrus.Fatalf("error running server :: %v", err)
			break
		}
	}

}
