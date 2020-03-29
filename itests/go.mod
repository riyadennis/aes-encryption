module github.com/riyadennis/aes-encryption/itests

go 1.13

require (
	github.com/riyadennis/aes-encryption/ex v0.0.0-20200328215420-464cc46dc941
	github.com/sirupsen/logrus v1.5.0
	google.golang.org/grpc v1.28.0
)

replace github.com/riyadennis/aes-encryption/ex => ../ex
