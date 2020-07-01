# configurator

配置器的特点:

1. 完整的配置管理功能
2. 多环境，多业务组，多服务，多版本 的配置分类

![architectures](https://raw.githubusercontent.com/xincao9/configurator/master/configurator.png)

## 安装中间件

**安装 [dkv](https://github.com/xincao9/dkv)**

```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest
```

**创建服务配置**

>接口
```
curl -X PUT -H 'content-type:application/json' 'http://localhost:9090/kv' -d '{"k":"configurator|TEST|BASE|USER-SERVICE|v1.0", "v":"{\"redis\":{\"host\":\"localhost\",\"port\":\"6379\"}}"}'
```

>推荐使用 [configurator-ui](https://github.com/xincao9/configurator/tree/master/api) 系统管理配置

## 如何使用JAVA SDK

[configurator-cli](https://github.com/xincao9/configurator-cli)

## 如何使用GO SDK

**获取SDK**

```
go get github.com/xincao9/configurator
```

**设置系统环境变量**

```
export env="TEST" // 环境
export group="BASE" // 业务组
export project="USER-SERVICE" // 项目
export version="v1.0" // 版本
export master="localhost:9090" // dkv 的master地址
export slaves="" // dkv 的slaves地址，host1:port1,host2:port2
```

**读取属性**

```
configurator.C.Get("redis.host")
configurator.C.Get("redis.port")
```

**设置管理端口**

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

**查看运行时的配置**

```
curl -X GET 'http://localhost:8080/config'
```

**知识**

* [dkv](https://github.com/xincao9/dkv)
* [viper](https://github.com/spf13/viper)
