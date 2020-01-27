GOFILES= $$(go list -f '{{join .GoFiles " "}}')

test:
	go test -timeout=5s -cover -race

run:
	go run $(GOFILES) server -config="config.yaml"

build:
	go build ./...

migrate:
	go run main.go -migrate=up

migrate_down:
	go run main.go -migrate=down

proto:
	protoc ex/api/data.proto --go_out=plugins=grpc:.

docker-build:
	docker build -t riyadennis/aes-encryption .

docker-run:
	docker run --rm -p 8086:8086  riyadennis/aes-encryption

docker-push:
	docker push riyadennis/aes-encryption:latest