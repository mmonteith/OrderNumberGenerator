DEFAULT_GOAL: all
VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`
BINARY=./build/distributions/native/order_number_generator
DOCKER_BINARY_LOCATION=./build/distributions/docker/
DOCKER_BINARY_NAME=order_number_generator

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

clean:
	@echo ${GOPATH}
	rm -Rf build
	if [ ! -d "vendor" ]; then \
		mkdir vendor; \
	else \
		rm -Rf vendor/*; \
	fi

dep:
	dep ensure

format:
	@echo "Build started at $(BUILD_TIME)"
	go fmt ./app/...

test:
	go test -cover -race ./...

all: $(SOURCES) clean dep format test
	go build -v -o ${BINARY} ./app/cmd/main.go


docker: $(SOURCES) clean dep format test
	GOOS=linux GOARCH=amd64 go build -v -o ${DOCKER_BINARY_LOCATION}${DOCKER_BINARY_NAME} app/cmd/main.go
	docker build -t docker.urbn.com/urbn/order_number_generator --build-arg DOCKER_BINARY_LOCATION=${DOCKER_BINARY_LOCATION} --build-arg DOCKER_BINARY_NAME=${DOCKER_BINARY_NAME} .
	docker images
	docker push docker.urbn.com/urbn/order_number_generator