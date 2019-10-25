all: eather application

dep:
	@echo "Install dependencies"
	go mod vendor

test:
	go test ./...

eather: dep
	@echo "Building eather console"
	go build -o eather \
		./console/main.go

application: dep
	@echo "Building eather app"
	go build -o eatherapp \
		./app/main.go
		
clean:
	go clean ./...
	rm -f eather eatherapp

.PHONY: eather application dep
