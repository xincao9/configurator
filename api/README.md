# configurator-ui

configurator ui Applicable to, need to support multi-environment, multi-version, multi-application centralized configuration


![UI](https://raw.githubusercontent.com/xincao9/configurator/master/api/resources/doc/configurator-ui.png)


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
conf: /usr/local/configurator-api/conf/config.yaml
```

**Acknowledgements**

* [dkv](https://github.com/xincao9/dkv)
