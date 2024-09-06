BINARY_NAME=my_redis

GO=go

BUILD_FLAGS=-ldflags "-s -w"

build:
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME) cmd/main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)