GOFILES= $$(go list -f '{{join .GoFiles " "}}')

clean:
	rm -rf vendor/

deps: clean
	glide install

test:
	go test -timeout=5s -cover -race $$(glide novendor)

run:
	go run $(GOFILES) server -config="config.yaml"

build:
	go build -o $(GOPATH)/bin/aes-encryption $(GOFILES)

migrate:
	go run main.go -migrate=up

migrate_down:
	go run main.go -migrate=down
proto:
	protoc ex/api/data.proto --go_out=plugins=grpc:.