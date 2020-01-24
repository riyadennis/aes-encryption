module github.com/riyadennis/aes-encryption

require (
	github.com/Microsoft/go-winio v0.4.12 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang-migrate/migrate v3.5.4+incompatible
	github.com/julienschmidt/httprouter v1.1.0
	github.com/onsi/ginkgo v1.5.0 // indirect
	github.com/onsi/gomega v1.4.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/riyadennis/aes-encryption/data v0.0.0-20200124133616-30a33ec1727f
	github.com/riyadennis/aes-encryption/ex v0.0.0-20191128221241-6769d6360d0f
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	google.golang.org/grpc v1.26.0
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

go 1.13

replace github.com/riyadennis/aes-encryption/data => ./data

replace github.com/riyadennis/aes-encryption/ex => ./ex
