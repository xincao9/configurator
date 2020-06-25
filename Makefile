test:
	go test -v . -cover

build:
	go build -o configurator-cli cmd/main.go

docker:
	docker build . -t configurator:latest

install:build
	mv configurator-cli /usr/local/bin/
