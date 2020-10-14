EXE := sprint
IMAGE := sprint
TARGET := distr
VERSION := $(shell git describe --tags --always --dirty)
TAG := $(VERSION)
REGISTRY := arpabet
PWD := $(shell pwd)
NOW := $(shell date +"%m-%d-%Y")

all: build

version:
	@echo $(TAG)

proto: version
	protoc proto/*.proto -I proto -I third_party -I $(GOPATH)/src/github.com/protocolbuffers/protobuf/src --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,allow_delete_body=true:. --swagger_out=logtostderr=true,allow_delete_body=true:.
	rm node.swagger.json
	mv *.swagger.json resources/swagger/

bindata: proto
	go-bindata -pkg resources -o pkg/resources/bindata.go -nocompress -nomemcopy -fs -prefix "resources/" resources/...
	go-bindata -pkg assets -o pkg/assets/bindata.go -nocompress -nomemcopy -fs -prefix "assets/" assets/...

build: bindata
	rm -rf rsrc.syso
	go test -cover ./...
	go build -o $(EXE) -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"

distr: build
	rm -rf $(TARGET)
	mkdir $(TARGET)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(TARGET)/$(EXE)_linux -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(TARGET)/$(EXE)_darwin -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(TARGET)/$(EXE).exe -v -ldflags "-X main.Version=$(VERSION) -X main.Build=$(NOW)"

run: build
	env COS=dev ./$(EXE)

test: build
	env COS=test ./$(EXE)

docker:
	docker build --build-arg VERSION=$(VERSION) --build-arg BUILD=$(NOW) -t $(REGISTRY)/$(IMAGE):$(TAG) -f Dockerfile .

docker-run: docker
	docker run -p 7000:7000 -p 7080:7080 --env COS $(REGISTRY)/$(IMAGE):$(TAG)

docker-build:
	rm -rf $(TARGET)/$(EXE)_linux
	docker build --build-arg VERSION=$(VERSION) --build-arg BUILD=$(NOW) -t $(REGISTRY)/$(IMAGE):$(TAG)-build -f Dockerfile.build .
	docker run --rm $(REGISTRY)/$(IMAGE):$(TAG)-build > $(TARGET)/$(EXE)_linux

docker-push: docker
	docker push ${REGISTRY}/${IMAGE}:${TAG}
	docker tag ${REGISTRY}/${IMAGE}:${TAG} ${REGISTRY}/${IMAGE}:latest
	docker push ${REGISTRY}/${IMAGE}:latest

clean:
	docker ps -q -f 'status=exited' | xargs docker rm
	echo "y" | docker system prune

licenses:
	go-licenses csv "github.com/arpabet/sprint" > resources/licenses.txt



