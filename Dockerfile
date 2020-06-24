FROM golang:1.12

LABEL maintainer="xincao9@gmail.com"

WORKDIR /go/src/configurator
COPY . .
RUN make install
