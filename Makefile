# Makefile

APPNAME=pig

all: build

build:
	go build -o $(APPNAME) main.go

clean:
	rm -f $(APPNAME)