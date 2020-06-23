# configurator

>Configuration Center Service

**Install dkv**

```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest

curl -X PUT -H 'content-type:application/json' 'http://localhost:9090/kv' -d '{"k":"configurator|test|cbs|user-service|1.0", "v":"{\"redis\":{\"host\":\"localhost\",\"port\":\"6379\"}}"}'
```

**System environment variables**

```
export env="test"
export group="cbs"
export project="user-service"
export version="1.0"
export master="localhost:9090"
export slaves=""
```

**Get the package**

```
go get github.com/xincao9/configurator
```

**Get configuration properties**

```
C.Get("redis.host")
C.Get("redis.port")
```

**Set the management endpoint**

```
http.HandleFunc("/config", configurator.AllSettings)
http.ListenAndServe(":8080", nil)
```

**View the current running configuration**

```
curl -X GET 'http://localhost:8080/config'
```
