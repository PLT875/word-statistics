NAME			= word-statistics
BINARY_PATH 	= ./$(NAME)
OS				= $(shell uname)

install:
	glide install

clean:
	go clean

unit_test:
	go test ./... -cover | grep -v vendor

build:
	go build -o $(BINARY_PATH)

run:
	$(BINARY_PATH) -p 5555