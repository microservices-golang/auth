LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml


build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/main.go	

copy-to-server:
	scp service_linux root@81.163.22.104:


docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 --provenance=false -t cr.selcloud.ru/microservices-golang/microservices-golang-server:v0.0.1 .
	docker login -u token -p CRgAAAAA8043s4cqSWPbuOugY1ZY0vzesKiu-IAB cr.selcloud.ru/microservices-golang
	docker push cr.selcloud.ru/microservices-golang/microservices-golang-server:v0.0.1