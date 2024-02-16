.PHONY: default
default: build

help: Makefile ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "; printf "Usage:\n\n    make \033[36m<target>\033[0m [VARIABLE=value...]\n\nTargets:\n\n"}; {printf "    \033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run app for development
	go run

.PHONY: test
test: ## Run tests
	go test -v --race -cpu 2 -cover -shuffle on -vet '' ./...

.PHONY: build
build:
	go build -ldflags="-s -w" -trimpath -o ./bin/pno

.PHONY: install
install: ## Install app
	install -b -S -v ./bin/pno ${GOPATH}/bin/.

.PHONY: clean
clean: ## Clean up
	rm -rf ./bin/*
	go mod tidy

.PHONY: release
release: ## Build release binaries
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o bin/pno-macos-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/pno-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -trimpath -o bin/pno-linux-386
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/pno-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -trimpath -o bin/pno-windows-386.exe
