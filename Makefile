.PHONY: default test build install clean release

default: build

run:
	go run cmd/genpno/main.go

test:
	go test -v --race -cpu 2 -cover -shuffle on -vet '' ./...

build:
	go build -o bin/genpno cmd/genpno/main.go

install:
	go install bin/genpno

clean:
	go mod tidy

release:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/genpno-macos-arm64 cmd/genpno/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/genpno-linux-amd64 cmd/genpno/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/genpno-linux-386 cmd/genpno/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/genpno-windows-amd64.exe cmd/genpno/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/genpno-windows-386.exe cmd/genpno/main.go
