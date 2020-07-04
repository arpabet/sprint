EXE := template-server
IMAGE := template-server
VERSION := $(shell git describe --tags --always --dirty)
TAG := $(VERSION)
REGISTRY := 0xcc
PWD := $(shell pwd)
NOW := $(shell date +"%m-%d-%Y")

all: build

version:
	@echo $(TAG)

bindata:
	go-bindata -pkg resources -o pkg/resources/bindata.go -nocompress -nomemcopy -prefix "resources/" resources/...

build: version
	go test -cover ./...
	go build  -v -ldflags "-X main.Version=$(VERSION) -X main.Built=$(NOW)"

run: build
	env COS=dev ./$(EXE)

test: build
	env COS=test ./$(EXE)

docker:
	docker build  --build-arg TAG=$(TAG) -t $(REGISTRY)/$(IMAGE):$(TAG) -f Dockerfile .

docker-run: docker
	docker run -p 8080:8080 -p 8081:8081 --env COS  $(REGISTRY)/$(IMAGE):$(TAG)

docker-push: docker
	docker push ${REGISTRY}/${IMAGE}:${TAG}
	docker tag ${REGISTRY}/${IMAGE}:${TAG} ${REGISTRY}/${IMAGE}:latest
	docker push ${REGISTRY}/${IMAGE}:latest

licenses:
	go-licenses csv "github.com/arpabet/template-server" > resources/licenses.txt