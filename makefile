BINARY_NAME=csv2json

test:
	go test convertor/*.go -v

build:
	cd src
	go build -o ${GOBIN}/${BINARY_NAME} main.go

run:
	go run src/main.go

clean:
	go clean
	rm ${GOBIN}/${BINARY_NAME}

deps:
    #go get github.com/gorilla/websocket

compile:
	echo "Compiling for every OS and Platform that matters"
	cd src
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${BINARY_NAME} main.go

all: deps build test compile
