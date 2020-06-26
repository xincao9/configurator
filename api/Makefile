test:
	go test -v . -cover

build:
	go build -o configurator-api main.go

docker:
	docker build . -t configurator-api:latest

install:build
	mkdir -p /usr/local/configurator/bin
	mv configurator-api /usr/local/configurator/bin/
	mkdir -p /usr/local/configurator/conf
	cp -rf resources/conf/* /usr/local/configurator/conf/
	mkdir -p /usr/local/configurator/assets
	cp -rf assets/* /usr/local/configurator/assets/
