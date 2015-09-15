build: all

cur-dir   := $(shell basename `pwd`)

all:
	go get github.com/fsouza/go-dockerclient
	go build main.go
	mv main docker-images-cleanup.bin

clean:
	rm docker-images-cleanup.bin