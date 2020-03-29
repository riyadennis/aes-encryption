package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/aes-encryption/internal/server"
)

func main() {
	server.Run()
}
