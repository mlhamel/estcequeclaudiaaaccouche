.SHELLFLAGS = -e
.NOTPARALLEL:

default: build

build:
	go build -o accouchement cmd/accouchement/main.go
