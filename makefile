BINARY_NAME=csv2json

test:
	go test convertor/*.go -v

build:
	go build -o ${GOBIN}/${BINARY_NAME} main.go

run:
	go run main.go

clean:
	go clean
	rm ${GOBIN}/${BINARY_NAME}

deps:
    #go get github.com/gorilla/websocket

compile:
	echo "Compiling for every OS and Platform that matters"
	GOOS=linux GOARCH=amd64 go build -o ${GOBIN}/${BINARY_NAME} main.go

all: deps build test compile