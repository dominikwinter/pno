.PHONY: default
default: build

.PHONY: run
run:
	go run

.PHONY: test
test:
	go test -v --race -cpu 2 -cover -shuffle on -vet '' ./...

.PHONY: build
build:
	go build -trimpath -o ./bin/pno

.PHONY: install
install:
	install -b -S -v ./bin/pno ${GOPATH}/bin/.

.PHONY: clean
clean:
	go mod tidy

.PHONY: release
release:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/pno-macos-arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pno-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/pno-linux-386
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/pno-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/pno-windows-386.exe
