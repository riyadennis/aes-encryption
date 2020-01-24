module github.com/riyadennis/aes-encryption/data

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/riyadennis/aes-encryption/ex v0.0.0-20191118220344-51b3c20d16fb
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/riyadennis/aes-encryption/ex => ../ex
