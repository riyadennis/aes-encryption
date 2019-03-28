package main
import (
	"github.com/riyadennis/aes-encryption/middleware"
	"github.com/riyadennis/aes-encryption/handlers"
	"flag"
	"github.com/riyadennis/aes-encryption/cmd"
	"os"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/sirupsen/logrus"
)
func main(){
	configFlag := flag.String("config", middleware.DefaultConfigPath, "Path to the config file")
	migrateFlag := flag.String("migrate", "up", "To Create tables up to delete them down ")
	flag.Parse()
	config, err := middleware.GetConfig(*configFlag)
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
}