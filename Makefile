test:
	go test -v . -cover

build:
	go build -o configurator-cli cmd/main.go
	go build -o configurator-api api/main.go

docker:
	docker build . -t configurator:latest

install:build
	mkdir -p /usr/local/configurator/bin
	mkdir -p /usr/local/configurator/conf
	mkdir -p /usr/local/configurator/assets
	mv configurator-cli /usr/local/configurator/bin
	mv configurator-api /usr/local/configurator/bin
