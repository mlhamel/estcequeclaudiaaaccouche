.SHELLFLAGS = -e
.NOTPARALLEL:

default: build

build:
		go install github.com/mlhamel/accouchement/web && go build -o bin/accouchement cmd/accouchement/main.go
