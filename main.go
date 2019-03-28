package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/riyadennis/aes-encryption/internal"
	"github.com/riyadennis/aes-encryption/internal/cmd"
	"github.com/riyadennis/aes-encryption/internal/handlers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/aes-encryption/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	configFlag := flag.String("config", internal.DefaultConfigPath, "Path to the config file")
	migrateFlag := flag.String("migrate", "up", "To Create tables up to delete them down ")
	flag.Parse()
	config, err := internal.GetConfig(*configFlag)
	if err != nil {
		logrus.Errorf("Unable to fetch config %s", err.Error())
	}
	if len(os.Args) < 2 {
		fmt.Println("Please enter an option")
		os.Exit(1)
	}
	if os.Args[1] == "-migrate=up" || os.Args[1] == "-migrate=down" {
		cmd.ExecuteCommand(*migrateFlag, config)
		os.Exit(0)
	}
	handlers.Run(config)
	server.Run()
}
