PROGNAME=httpfromtcp
.PHONY: build run

build:
	go build -o bin/$(PROGNAME)

run: build
	./bin/$(PROGNAME)