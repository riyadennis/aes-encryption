package main

import (
	"flag"

	"github.com/riyadennis/aes-encryption/ex"
	"github.com/riyadennis/aes-encryption/internal/handlers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/aes-encryption/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	configFlag := flag.String("config", ex.DefaultConfigPath, "Path to the config file")
	flag.Parse()
	config, err := ex.GetConfig(*configFlag)
	if err != nil {
		logrus.Errorf("Unable to fetch config %s", err.Error())
	}
	handlers.Run(config)
	server.Run()
}
