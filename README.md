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
**Set configuration properties**

```
caoxindeMacBook-Air:configurator caoxin$ go run github.com/xincao9/configurator/cmd/main.go
Usage of /var/folders/7y/0qk9vb817r7c9kr8_bjk7kym0000gn/T/go-build064844085/b001/exe/main:
  -env string
    	environment
  -group string
    	group
  -master string
    	dkv master address (default "localhost:9090")
  -project string
    	project
  -properties string
    	configuration properties
  -version string
    	version
```

**Get configuration properties**

```
configurator.C.Get("redis.host")
configurator.C.Get("redis.port")
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

**Acknowledgements**

* [dkv](https://github.com/xincao9/dkv)
* [viper](https://github.com/spf13/viper)
