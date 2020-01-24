module github.com/riyadennis/aes-encryption/ex

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/pkg/errors v0.9.1
	github.com/riyadennis/aes-encryption/data v0.0.0-20200124133616-30a33ec1727f
	github.com/stretchr/testify v1.4.0
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/riyadennis/aes-encryption/data => ../data
