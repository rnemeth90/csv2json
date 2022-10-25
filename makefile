BINARY_NAME=csv2json

test:
	go test -v main.go

build:
	go build -o ${BINARY_NAME} main.go

run:
	go run main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

deps:
    #go get github.com/gorilla/websocket

compile:
	echo "Compiling for every OS and Platform that matters"
	GOOS=linux GOARCH=arm go build -o ${GOPATH}/bin/${BINARY_NAME}-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o ${GOPATH}/bin/${BINARY_NAME}-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o ${GOPATH}/bin/${BINARY_NAME}-linux-amd64 main.go

all: deps build test