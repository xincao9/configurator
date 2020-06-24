test:
	go test -v . -cover

build:
	go build -o configurator-api main.go

docker:
	docker build . -t configurator-api:latest

install:build
	mkdir -p /usr/local/configurator-api/bin
	mv configurator-api /usr/local/configurator-api/bin/
	mkdir -p /usr/local/configurator-api/conf
	cp -rf resources/conf/* /usr/local/configurator-api/conf/
	mkdir -p /usr/local/configurator-api/assets
	cp -rf assets/* /usr/local/configurator-api/assets/
