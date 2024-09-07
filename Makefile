BINARY_NAME=my_redis

GO=go

BUILD_FLAGS=-ldflags "-s -w"

build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME) cmd/main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
	rm -f coverage.out
	rm -f coverage.html

test:
	$(GO) test -v ./...

coverage:
	$(GO) test -coverprofile=coverage.out ./...
	make gen_coverage_report

gen_coverage_report :
	$(GO) tool cover -html=coverage.out -o coverage.html