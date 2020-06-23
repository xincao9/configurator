# configurator

>Configure Central Services

**install dkv**
```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest
```

**system env variable**
```
export env="test"
export group="cbs"
export project="user-service"
export version="1.0"
export master="localhost:9090"
export slaves=""
```

**sdk**

```
go get github.com/xincao9/configurator

configurator.C.Get ("key") return interface{}
```
