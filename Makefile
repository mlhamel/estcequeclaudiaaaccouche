.SHELLFLAGS = -e
.NOTPARALLEL:

default: build

build:
		go build -o bin/accouchement
