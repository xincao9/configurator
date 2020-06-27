# configurator-ui

配置器的特点:

1. 完整的配置管理功能
2. 多环境，多业务组，多服务，多版本 的配置分类

## 安装中间件

**安装 [dkv](https://github.com/xincao9/dkv)**

```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest
```

**MYSQL**

* [schema](https://github.com/xincao9/configurator/blob/master/api/resources/doc/schema.sql)
* [data](https://github.com/xincao9/configurator/blob/master/api/resources/doc/data.sql)
* username/password: admin/admin

**安装 configurator-ui**

```
git clone https://github.com/xincao9/configurator.git
cd configurator/api
sudo make install
可执行文件 /usr/local/configurator/bin/configurator-api
配置文件: /usr/local/configurator/conf/configurator-api.yaml
```

**configurator-api.yaml**

```
db:
    datasourcename: root:asdf@tcp(localhost:3306)/configurator?charset=utf8&parseTime=true
dkv:
    address: localhost:9090
logger:
    dir: "/tmp/logs"
    level: debug
manager:
    server:
        port: 8090
server:
    cors:
        accesscontrolalloworigin: http://localhost:8081
    mode: debug
    port: 8080
```

**知识**

* [dkv](https://github.com/xincao9/dkv)
