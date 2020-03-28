module github.com/riyadennis/aes-encryption

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/julienschmidt/httprouter v1.1.0
	github.com/riyadennis/aes-encryption/data v0.0.0-20200124133616-30a33ec1727f
	github.com/riyadennis/aes-encryption/ex v0.0.0-20191128221241-6769d6360d0f
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	google.golang.org/grpc v1.27.1
)

go 1.13

replace github.com/riyadennis/aes-encryption/data => ./data

replace github.com/riyadennis/aes-encryption/ex => ./ex
