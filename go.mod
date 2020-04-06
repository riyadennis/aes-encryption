module github.com/riyadennis/aes-encryption

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/riyadennis/aes-encryption/ex v0.0.0-20191128221241-6769d6360d0f
	github.com/sirupsen/logrus v1.5.0
	go.mongodb.org/mongo-driver v1.3.1
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	google.golang.org/grpc v1.27.1
)

go 1.13

replace github.com/riyadennis/aes-encryption/ex => ./ex
