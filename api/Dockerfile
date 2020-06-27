FROM golang:1.12

LABEL maintainer="xincao9@gmail.com"

WORKDIR /go/src/configurator/api
COPY . .
RUN make install

EXPOSE 8080

CMD ["/usr/local/configurator/bin/configurator-api"]
