.SHELLFLAGS = -e
.NOTPARALLEL:

default: build

build:
  go install ./...
	go build -o bin/accouchement cmd/accouchement/main.go
