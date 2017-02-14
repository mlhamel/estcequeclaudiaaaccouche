.SHELLFLAGS = -e
.NOTPARALLEL:

default: build

build:
		go build -o bin/accouchement cmd/accouchement/main.go
