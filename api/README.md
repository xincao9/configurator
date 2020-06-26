# configurator-ui


>configurator ui Applicable to, need to support multi-environment, multi-version, multi-application centralized configuration

**Install dkv**

```
docker pull xincao9/dkv
docker run -d -p 9090:9090 -p 6380:6380 dkv:latest
```

**Install configurator-ui**

```
git clone https://github.com/xincao9/configurator.git

cd configurator/api

sudo make install

bin: /usr/local/configurator-api/bin/configurator-api
conf: /usr/local/configurator-api/conf/configurator-api.yaml
```

**Configuration file**

* [schema](https://github.com/xincao9/configurator/blob/master/api/resources/doc/schema.sql)
* [data](https://github.com/xincao9/configurator/blob/master/api/resources/doc/data.sql)
* username/password: admin/admin
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

![UI](https://raw.githubusercontent.com/xincao9/configurator/master/api/resources/doc/configurator-ui.png)

**Acknowledgements**

* [dkv](https://github.com/xincao9/dkv)
