---
author: "jdlau"
date: 2021-12-09
linktitle: docker compose使用extra host让容器访问主机服务
menu:
next:
prev:
title: docker compose使用extra host让容器访问主机服务
weight: 10
---

[首发于：简单博客](/posts/2021/12/docker_compose_extra_host/)

## docker compose 如何访问主机服务

docker compose 里面的容器怎么访问主机自身起的服务呢？

[20.10.0 版本在 linux 新增 host.docker.internal 支持](https://docs.docker.com/engine/release-notes/#networking-3)：
`docker run -it --add-host=host.docker.internal:host-gateway alpine cat /etc/hosts`

```sh
127.0.0.1       localhost
::1     localhost ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
172.17.0.1      host.docker.internal # --add-host的作用就是添加了这行到/etc/hosts
172.17.0.3      cb0565ceea26
```

[相关提交](https://github.com/moby/moby/pull/40007)

这个 add-host 的意思是告诉容器，容器对域名 host.docker.internal 的访问都将转发到 host-gateway 去。

也就是容器内部访问这个域名 host.docker.internal 时，就会访问到对应的主机上的 host-gateway 地址。

从而达到容器访问主机上服务的效果。

那么，这个 add-host 怎么用在 compose 上呢？

[在 build 里使用 extra_hosts](https://github.com/docker/cli/issues/1293)

```yaml
version: "2.3" # 因为某个bug的存在，只能用version2，不能用version3
services:
  tmp:
    build:
      context: .
    extra_hosts: # 配置extra_hosts
      - "host:IP"
    command: -kIL https://host
    tty: true
    stdin_open: true
```

[docker compose 配置中文说明](https://www.huaweicloud.com/articles/d8c4873d55e2485840070b65765860b9.html)

[参照](https://docs.microsoft.com/en-us/dotnet/architecture/microservices/multi-container-microservice-net-applications/multi-container-applications-docker-compose)

测试：

```sh
docker-compose --version
docker-compose version 1.29.2, build 5becea4c
```

新建一个服务，在主机上运行；

```go
package main

import (
        "fmt"
        "net/http"
)

func main() {
        handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
                fmt.Println("hi")
                resp.Write([]byte("hello"))
        })
        if err := http.ListenAndServe(":8080", handler); err != nil {
                panic(err)
        }
}
```

新建 compose，里面也起一个服务，这个服务需要访问上述的主机服务；

```yaml
version: "2.3" # version改为3.3也可以
services:
  server:
    image: curlimages/curl
    command: curl http://host.docker.internal:8080
    extra_hosts:
      - "host.docker.internal:host-gateway"
```

在终端访问容器服务，容器服务访问主机服务，如果能正常执行，则表示完成。

执行`docker-compose up`，能看到请求成功。

[代码](https://github.com/donnol/compose)
