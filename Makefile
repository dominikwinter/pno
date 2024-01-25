.PHONY: default b build i install r release

default: build

b build:
	go build

i install:
	go install

r release:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/genpno-macos-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/genpno-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/genpno-linux-386
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/genpno-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/genpno-windows-386.exe
