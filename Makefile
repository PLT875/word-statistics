NAME             = word-statistics
BINARY_PATH      = ./$(NAME)
OS               = $(shell uname)
# PACKAGES       = `glide novendor | grep -v "features" | grep -v "^\.$$"` | grep -v "fakes"

install:
	glide install

clean:
	go clean

unit_test:
	go test -cover ./aggregator/... ./api/... .
	# go test -cover $(PACKAGES)

build:
	go build -o $(BINARY_PATH)

run:
	$(BINARY_PATH) -ingest 5555 -stats 8080