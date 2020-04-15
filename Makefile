GOFILES= $$(go list -f '{{join .GoFiles " "}}')

test:
	go test -timeout=5s -cover -race

run:
	go run $(GOFILES) server -config="config.yaml"

build:
	go build ./...

proto:
	protoc ex/api/data.proto --go_out=plugins=grpc:.

docker-build:
	docker build -t riyadennis/aes-encryption .

docker-run:
	docker run --rm -p 8082:8082  riyadennis/aes-encryption

docker-push:
	docker push riyadennis/aes-encryption:latest