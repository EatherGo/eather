all: eahter

dep:
	@echo "Install dependencies"
	go mod vendor

test:
	go test ./...

eahter:
	@echo "Building eather console"
	go build ./eather/main.go
	cp
		
client: protoc
	@echo "Building client"
	go build -o client \
		./src/client

clean:
	go clean ./...
	rm -f server client

.PHONY: client server protoc dep
