GOPATH:=$(shell go env GOPATH)

test:
	go test -v . -cover

