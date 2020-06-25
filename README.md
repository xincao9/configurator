# configurator

>Configuration Center Service Applicable to, need to support multi-environment, multi-version, multi-application centralized configuration

![architectures](https://raw.githubusercontent.com/xincao9/configurator/master/configurator.png)

**Install dkv**

```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest

curl -X PUT -H 'content-type:application/json' 'http://localhost:9090/kv' -d '{"k":"configurator|test|cbs|user-service|1.0", "v":"{\"redis\":{\"host\":\"localhost\",\"port\":\"6379\"}}"}'
```

**Get the package**

```
go get github.com/xincao9/configurator
```
**Set configuration properties**

It is recommended to use configurator-ui to manage configuration

```
#The following is set through the command line

install: make install
exec: configurator-cli
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


**System environment variables**

```
export env="test"
export group="cbs"
export project="user-service"
export version="1.0"
export master="localhost:9090"
export slaves=""
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

```
r := gin.Default()
r.GET("/config", func(c *gin.Context) {
    c.JSON(http.StatusOK, configurator.C.AllSettings())
})
```

**View the current running configuration**

```
curl -X GET 'http://localhost:8080/config'
```

**Acknowledgements**

* [dkv](https://github.com/xincao9/dkv)
* [viper](https://github.com/spf13/viper)
