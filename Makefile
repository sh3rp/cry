.PHONY: all install

all: install

install:
	dep ensure
	go install cmd/cry/cry.go