---
author: "jdlau"
date: 2021-12-10
linktitle: k8s
menu:
next:
prev:
title: k8s
weight: 10
categories: ['k8s']
tags: ['k8s']
---

## What

docker 带来容器之风，以致容器多不胜数。如何编排和管理众多容器，使得它们同心协力办好事情，即成为了当下最大的课题。

为此，k8s 应运而生。

容器，通讯，存储，配置。

## Why

为编排和管理数量众多的容器。

## How

### Install

#### k8s: 集群搭建所需资源

> One or more machines running one of:
>
> > Ubuntu 16.04+
> >
> > Debian 9+
> >
> > CentOS 7+
> >
> > Red Hat Enterprise Linux (RHEL) 7+
> >
> > Fedora 25+
> >
> > HypriotOS v1.0.1+
> >
> > Flatcar Container Linux (tested with 2512.3.0)
>
> 2 GB or more of RAM per machine (any less will leave little room for your apps).
>
> 2 CPUs or more.
>
> Full network connectivity between all machines in the cluster (public or private network is fine).
>
> Unique hostname, MAC address, and product_uuid for every node. See here for more details.
>
> Certain ports are open on your machines. See here for more details.
>
> Swap disabled. You MUST disable swap in order for the kubelet to work properly.

[参考官方文档](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#before-you-begin)

#### 本地机器起多个虚拟机搭建

#### 使用 virtualBox 创建三台虚拟机

1.用 NAT 和 host only 网络模式

> virtualBox 安装比较简单，不再介绍，GUI 工具用起来也很方便，这部分只介绍我认为需要提示的部分。
>
> **内存推荐 2048M, CPU 推荐 2 个**
>
> 默认只有一个 NAT 适配器，添加一个 **Host-Only Adapter**。NAT 适配器是虚拟机用来访问互联网的，Host-Only 适配器是用来虚拟机之间通信的。
>
> 以 Normal Start 方式启动虚拟机安装完系统以后，因为是 server 版镜像，所以没有图形界面，直接使用用户名密码登录即可。
>
> 修改配置，**enp0s8 使用静态 IP**。配置请参考 SSH between Mac OS X host and Virtual Box guest。注意配置时将其中的网络接口名改成你自己的 Host-Only Adapter 对应的接口。
>
> 一台虚拟机创建完成以后可以使用 clone 方法复制出两台节点出来，注意 clone 时为新机器的网卡重新初始化 MAC 地址。
>
> 三台虚拟机的静态 IP 都配置好以后就可以使用 ssh 在本地主机的终端上操作三台虚机了。虚机使用 Headless Start 模式启动

[参照](https://github.com/c-rainstorm/blog/blob/master/devops/%E6%9C%AC%E6%9C%BA%E6%90%AD%E5%BB%BA%E4%B8%89%E8%8A%82%E7%82%B9k8s%E9%9B%86%E7%BE%A4.md)

[参照 2](<http://www.zchengjoey.com/posts/Ubuntu1604%E6%90%AD%E5%BB%BAk8s%E9%9B%86%E7%BE%A4(%E9%99%84%E5%B8%A6docker%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8%E4%BB%A3%E7%90%86)/>)

2.用桥接网络模式

vm: 虚拟机上安装 k8s

先添加 key：`https_proxy=http://192.168.56.1:51837 curl -s -v https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -`

然后添加 source：

```sh
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
```

最后更新：`sudo apt -o Acquire::http::proxy="http://192.168.56.1:51837" update`

kubeadm

安装：`sudo apt -o Acquire::https::proxy="http://192.168.56.1:51837" install -y kubeadm`

kubelet

安装：`sudo apt -o Acquire::https::proxy="http://192.168.56.1:51837" install -y kubelet`

kubectl

安装：`sudo apt -o Acquire::https::proxy="http://192.168.56.1:51837" install -y kubectl`

docker

[参照](https://kubernetes.io/zh/docs/setup/production-environment/container-runtimes/#docker)

从上面的参照可以看到除了 docker 之外，还可以选择其它运行时。

k8s: 三个工具

kubeadm

作用：用来**初始化集群**的指令。

使用：

> 在一台机器上执行`kubeadm init`，初始化集群，该机器作为集群 master；初始化成功后会返回`<arguments-returned-from-init>`，作为其它机器加入该集群的参数 。
>
> 在另一台机器上执行`kubeadm join <arguments-returned-from-init>`。
>
> 如果想添加更多机器，请重复`kubeadm join`指令。

kubelet

作用：在集群中的每个节点上用来**启动 Pod 和容器**等。

在每个节点上运行的节点代理。它可以向 apiserver 注册节点。

> The kubelet works in terms of a PodSpec. A PodSpec is a YAML or JSON object
> that describes a pod.
>
> -- kubelet 在一系列 pod 规范里工作。一个 pod 规范是一个描述 pod 的 yaml 或 json 对象。

kubectl

作用：用来**与集群通信**的命令行工具。

创建资源，暴露服务。

[参照](https://kubernetes.io/zh/docs/setup/production-environment/tools/kubeadm/install-kubeadm/)

### k8s: 初始化集群

#### 关闭 swap 交换空间

执行命令：`sudo swapoff -a`，并将文件`/etc/fstab`里关于 swap 的行注释掉，然后重启机器

为什么要关闭呢？

[issue 上的讨论 1](https://github.com/kubernetes/kubernetes/issues/53533)

[issue 上的讨论 2](https://github.com/kubernetes/kubernetes/issues/7294)

> having swap available has very strange and bad interactions with memory limits
>
> kubernetes 的想法是将实例紧密包装到尽可能接近 100％。 所有的部署应该与 CPU /内存限制固定在一起。 所以如果调度程序发送一个 pod 到一台机器，它不应该使用交换。 设计者不想交换，因为它会减慢速度。
>
> 所以关闭 swap 主要是为了性能考虑。
>
> 当然为了一些节省资源的场景，比如运行容器数量较多，可添加 kubelet 参数 --fail-swap-on=false 来解决。

#### 初始化

先通过命令`sudo kubeadm config print init-defaults > init-default.yaml`生成默认配置文件。

> 在生成的 yaml 配置文件里修改:
>
> `advertiseAddress: 192.168.9.43`，其中的 ip 地址为`ip a`拿到的地址。
>
> networking:
> dnsDomain: cluster.local
> podSubnet: 10.46.128.0/21 # 这个 ip 将要用在安装成功后的 pod network 配置里。
> serviceSubnet: 192.168.1.0/24

然后在初始化命令使用该配置文件：`sudo kubeadm init --config=init-default.yaml --v=5`

出现警告：

> [preflight] Running pre-flight checks
>
> [WARNING IsDockerSystemdCheck]: detected "cgroupfs" as the Docker cgroup driver. The recommended driver is "systemd". Please follow the guide at <https://kubernetes.io/docs/setup/cri/>

里面说到 cgroup 驱动用了 cgroupfs，而不是 systemd。可参照[官方文档](https://kubernetes.io/docs/setup/cri/)修改设置。

然后会到`k8s.gcr.io`获取镜像，这时又出现超时错误。应该是墙导致的，需要使用代理或[改用镜像站](https://vqiu.cn/how-to-access-gcr-io/)。

[镜像站也关掉了，怎么办？](https://developer.aliyun.com/article/759310)

使用 docker 拉取镜像，然后修改 tag：

> 先看所需镜像：`sudo kubeadm config images list`
>
> > k8s.gcr.io/kube-apiserver:v1.20.3
> >
> > k8s.gcr.io/kube-controller-manager:v1.20.3
> >
> > k8s.gcr.io/kube-scheduler:v1.20.3
> >
> > k8s.gcr.io/kube-proxy:v1.20.3
> >
> > k8s.gcr.io/pause:3.2
> >
> > k8s.gcr.io/etcd:3.4.13-0
> >
> > k8s.gcr.io/coredns:1.7.0
>
> 逐个从 docker hub 上搜到并获取镜像：`sudo docker pull aiotceo/kube-apiserver:v1.20.3`
>
> > sudo docker pull aiotceo/kube-apiserver:v1.20.3
> >
> > sudo docker pull aiotceo/kube-controller-manager:v1.20.3
> >
> > sudo docker pull aiotceo/kube-scheduler:v1.20.3
> >
> > sudo docker pull aiotceo/kube-proxy:v1.20.3
> >
> > sudo docker pull aiotceo/pause:3.2
> >
> > sudo docker pull bitnami/etcd:3.4.13
> >
> > sudo docker pull aiotceo/coredns:1.7.0
>
> 修改 tag：`sudo docker tag aiotceo/kube-apiserver:v1.20.3 k8s.gcr.io/kube-apiserver:v1.20.3`
>
> 最后删除：`sudo docker rmi aiotceo/kube-apiserver:v1.20.3`
>
> [preflight] You can also perform this action in beforehand using 'kubeadm config images pull'
>
> -- 拉取镜像也可以提前使用`kubeadm config images pull`完成。

kubelet 未启动错误：

> [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
> I0218 09:29:21.630514 98836 request.go:943] Got a Retry-After 1s response for attempt 1 to <https://10.0.2.15:6443/healthz?timeout=10s> > [kubelet-check] Initial timeout of 40s passed.
> I0218 09:29:43.755024 98836 request.go:943] Got a Retry-After 1s response for attempt 1 to <https://10.0.2.15:6443/healthz?timeout=10s>
> I0218 09:30:21.758311 98836 request.go:943] Got a Retry-After 1s response for attempt 1 to <https://10.0.2.15:6443/healthz?timeout=10s>
> I0218 09:32:24.612682 98836 request.go:943] Got a Retry-After 1s response for attempt 1 to <https://10.0.2.15:6443/healthz?timeout=10s>

```sh
        Unfortunately, an error has occurred:
                timed out waiting for the condition

        This error is likely caused by:
                - The kubelet is not running
                - The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)

        If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
                - 'systemctl status kubelet'
                - 'journalctl -xeu kubelet'

        Additionally, a control plane component may have crashed or exited when started by the container runtime.
        To troubleshoot, list all containers using your preferred container runtimes CLI.

        Here is one example how you may list all Kubernetes containers running in docker:
                - 'docker ps -a | grep kube | grep -v pause'
                Once you have found the failing container, you can inspect its logs with:
                - 'docker logs CONTAINERID'
```

[原来是在关闭了 swap 后要重启](https://stackoverflow.com/questions/52119985/kubeadm-init-shows-kubelet-isnt-running-or-healthy)

初始化过程中报错了之后需要重置一下才行：`sudo kubeadm reset --v=5`

否则会报错：

> [preflight] Some fatal errors occurred:

```sh
[ERROR Port-10259]: Port 10259 is in use
[ERROR Port-10257]: Port 10257 is in use
[ERROR FileAvailable--etc-kubernetes-manifests-kube-apiserver.yaml]: /etc/kubernetes/manifests/kube-apiserver.yaml already exists
[ERROR FileAvailable--etc-kubernetes-manifests-kube-controller-manager.yaml]: /etc/kubernetes/manifests/kube-controller-manager.yaml already exists
[ERROR FileAvailable--etc-kubernetes-manifests-kube-scheduler.yaml]: /etc/kubernetes/manifests/kube-scheduler.yaml already exists
[ERROR FileAvailable--etc-kubernetes-manifests-etcd.yaml]: /etc/kubernetes/manifests/etcd.yaml already exists
[ERROR Port-10250]: Port 10250 is in use

[preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
error execution phase preflight
```

查看`kubelet`日志：`journalctl -xeu kubelet`

`Failed to connect to 10.0.2.15 port 6443: Connection refused`:

[除了防火墙和 swap，还要关闭 selinux](https://zhuanlan.zhihu.com/p/265968760)

> 查看防火墙状态：systemctl status firewalld
>
> 临时关闭 selinux：
>
> > `sudo setenforce 0`
>
> 永久关闭：
>
> > 执行`sestatus`查看 selinux 状态，需要先安装工具`sudo apt install policycoreutils`
> >
> > 查看后，再去编辑配置文件`sudo vim /etc/selinux/config`，改为`SELINUX=disabled`

[需要翻墙的情况下](https://zhuanlan.zhihu.com/p/31398416)

#### 修改 docker 的 cgroup driver 为 systemd

为了确保 docker 和 kubelet 的 cgroup driver 一样，需要将 docker 的 cgroup driver 改为 systemd。

查看：`sudo docker info`，发现`Cgroup Driver: cgroupfs`

修改：`sudo vim /etc/docker/daemon.json`，添加以下内容：

```json
{
  "exec-opts": ["native.cgroupdriver=systemd"]
}
```

重启：`sudo systemctl daemon-reload` `sudo systemctl restart docker`

再次查看 docker info，会看到`Cgroup Driver: systemd`

#### kubectl

执行`sudo kubelet`时出现错误：

> failed to get the kubelet's cgroup: cpu and memory cgroup hierarchy not unified. cpu: /user.slice, memory: /user.slice/user-1000.slice/session-1.scope. Kubelet system container metrics may be missing.

难道是不能这样直接执行，需要用 systemctl 来执行：`sudo systemctl start kubelet.service`

通过查看`sudo cat /etc/systemd/system/kubelet.service.d/10-kubeadm.conf`配置文件，发现里面有很多 env 设置。

在`/etc/systemd/system/kubelet.service.d/10-kubeadm.conf`配置文件里添加`--cgroup-driver=systemd`配置。

是不是虚拟机创建时选择的网卡错了（用了 NAT 和 host only），导致上面的这么多问题呢？

[换成 NAT network 或者桥接试下](https://www.cnblogs.com/woncode/p/12206023.html)。

> With bridged networking, Oracle VM VirtualBox uses a device driver on your host system that filters data from your physical network adapter. This driver is therefore called a net filter driver. This enables Oracle VM VirtualBox to intercept data from the physical network and inject data into it, effectively creating a new network interface in software. When a guest is using such a new software interface, it looks to the host system as though the guest were physically connected to the interface using a network cable. The host can send data to the guest through that interface and receive data from it. This means that you can set up routing or bridging between the guest and the rest of your network.

[vbox 官方文档关于网卡的介绍](https://www.virtualbox.org/manual/ch06.html)

#### 成功

成功后会有以下信息：

其中需要注意部分信息是需要执行的命令，还有其它节点添加到集群时需要用到的信息。

```sh
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.9.43:6443 --token abcdef.0123456789abcdef \
    --discovery-token-ca-cert-hash sha256:5660805073db31916952821a8751ca0ee0644ce4205f616805f8a7f175ff8b33
```

[添加 pod network：](https://github.com/flannel-io/flannel#deploying-flannel-manually)

> 下载配置文件：wget <https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml>
>
> 编辑配置：修改
>
> net-conf.json: |
>
> {
> "Network": "10.46.128.0/21", # 这个地址要与 kubeadm 指定的配置文件里的 podSubnet 里的 ip 一致
> "Backend": {
> "Type": "vxlan"
> }
> }
>
> 最后执行命令: `kubectl apply -f kube-flannel.yml`

最后，查看 pods 状态：`kubectl get pods --all-namespaces`

```sh
jd@jnode:~$ kubectl get pods --all-namespaces
NAMESPACE     NAME                            READY   STATUS    RESTARTS   AGE
kube-system   coredns-74ff55c5b-f76bb         1/1     Running   0          5m52s
kube-system   coredns-74ff55c5b-frgkw         1/1     Running   0          5m52s
kube-system   etcd-jnode                      1/1     Running   0          6m8s
kube-system   kube-apiserver-jnode            1/1     Running   0          6m8s
kube-system   kube-controller-manager-jnode   1/1     Running   0          6m8s
kube-system   kube-flannel-ds-xghcx           1/1     Running   0          4m4s
kube-system   kube-proxy-nnbbx                1/1     Running   0          5m52s
kube-system   kube-scheduler-jnode            1/1     Running   0          6m8s
```

全部正常运行，完成。

#### 错误

k8s 的安装过程，可谓是一波三折。下面整理一下遇到的错误：

[Error registering network: failed to acquire lease: node "nodeName" pod cidr not assigned
](https://github.com/flannel-io/flannel/issues/1344)

> 确保 Network 的值与 podSubnet 一致即可。

#### 添加 worker 节点

在初始化完成后的输出信息里有：

```sh
Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.9.43:6443 --token abcdef.0123456789abcdef \
    --discovery-token-ca-cert-hash sha256:5660805073db31916952821a8751ca0ee0644ce4205f616805f8a7f175ff8b33
```

这个命令就是 worker 要加入集群时需要执行的命令。

在执行上述命令之前，需要[重新生成 token 和 hash](https://cloud.tencent.com/developer/news/268376):

```sh
# 重新生成新的token
kubeadm token create
kubeadm token list

# 获取ca证书sha256编码hash值，拿到的hash值要记录下来
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'

# 节点加入集群，token和hash值使用上面生成的
kubeadm join ip:6443 --token xxx --discovery-token-ca-cert-hash sha256:xxx
```

最后在主节点，执行：`kubectl get nodes`查看已加入的节点。

重启之后要等一会才会正常：

```sh
$ kubectl get nodes
The connection to the server 192.168.9.43:6443 was refused - did you specify the right host or port?
```

如果没有执行`export KUBECONFIG=/etc/kubernetes/admin.conf`，那么在使用 kubectl 命令时就不能在前面加 sudo 了。

在主节点成功看到 worker node 之后，还需要将主节点里的`~/.kube/config`文件复制到 worker node 上：

```sh
# 在主节点上使用scp将配置文件复制到worker
scp ~/.kube/config xx@xxx.xxx.xxx.xxx:~/.kube/config
```

否则会报错：

```sh
kubectl describe node
The connection to the server localhost:8080 was refused - did you specify the right host or port?
```

复制配置到 worker 后，在 worker 节点上执行`kubectl describe node`，就能看到节点信息。

[参照](https://stackoverflow.com/questions/63539796/connection-refused-error-on-worker-node-in-kubernetes)

#### k8s: 初始化配置文件

##### init-default.yaml

```yaml
apiVersion: kubeadm.k8s.io/v1beta2
bootstrapTokens:
  - groups:
      - system:bootstrappers:kubeadm:default-node-token
    token: abcdef.0123456789abcdef
    ttl: 24h0m0s
    usages:
      - signing
      - authentication
kind: InitConfiguration
localAPIEndpoint:
  advertiseAddress: 192.168.9.43 # 本机的ip地址，可通过`ip a`获取
  bindPort: 6443
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: jnode
  taints:
    - effect: NoSchedule
      key: node-role.kubernetes.io/master
---
apiServer:
  timeoutForControlPlane: 4m0s
apiVersion: kubeadm.k8s.io/v1beta2
certificatesDir: /etc/kubernetes/pki
clusterName: kubernetes
controllerManager: {}
dns:
  type: CoreDNS
etcd:
  local:
    dataDir: /var/lib/etcd
imageRepository: k8s.gcr.io
kind: ClusterConfiguration
kubernetesVersion: v1.20.0
networking:
  dnsDomain: cluster.local
  podSubnet: 10.46.128.0/21 # pod subnet值需要与kube-flannel.yml里的网络设置一致
  serviceSubnet: 192.168.1.0/24 # 跟本机ip同一网段
scheduler: {}
```

## kube-flannel.yml

```yaml
---
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: psp.flannel.unprivileged
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: docker/default
    seccomp.security.alpha.kubernetes.io/defaultProfileName: docker/default
    apparmor.security.beta.kubernetes.io/allowedProfileNames: runtime/default
    apparmor.security.beta.kubernetes.io/defaultProfileName: runtime/default
spec:
  privileged: false
  volumes:
    - configMap
    - secret
    - emptyDir
    - hostPath
  allowedHostPaths:
    - pathPrefix: "/etc/cni/net.d"
    - pathPrefix: "/etc/kube-flannel"
    - pathPrefix: "/run/flannel"
  readOnlyRootFilesystem: false
  # Users and groups
  runAsUser:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
  # Privilege Escalation
  allowPrivilegeEscalation: false
  defaultAllowPrivilegeEscalation: false
  # Capabilities
  allowedCapabilities: ["NET_ADMIN", "NET_RAW"]
  defaultAddCapabilities: []
  requiredDropCapabilities: []
  # Host namespaces
  hostPID: false
  hostIPC: false
  hostNetwork: true
  hostPorts:
    - min: 0
      max: 65535
  # SELinux
  seLinux:
    # SELinux is unused in CaaSP
    rule: "RunAsAny"
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
rules:
  - apiGroups: ["extensions"]
    resources: ["podsecuritypolicies"]
    verbs: ["use"]
    resourceNames: ["psp.flannel.unprivileged"]
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
  - apiGroups:
      - ""
    resources:
      - nodes
    verbs:
      - list
      - watch
  - apiGroups:
      - ""
    resources:
      - nodes/status
    verbs:
      - patch
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: flannel
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: flannel
subjects:
  - kind: ServiceAccount
    name: flannel
    namespace: kube-system
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: flannel
  namespace: kube-system
---
kind: ConfigMap
apiVersion: v1
metadata:
  name: kube-flannel-cfg
  namespace: kube-system
  labels:
    tier: node
    app: flannel
data:
  cni-conf.json: |
    {
      "name": "cbr0",
      "cniVersion": "0.3.1",
      "plugins": [
        {
          "type": "flannel",
          "delegate": {
            "hairpinMode": true,
            "isDefaultGateway": true
          }
        },
        {
          "type": "portmap",
          "capabilities": {
            "portMappings": true
          }
        }
      ]
    }
  net-conf.json: |
    {
      "Network": "10.46.128.0/21", # 与init-default.yaml配置文件里的podSubnet值一致
      "Backend": {
        "Type": "vxlan"
      }
    }
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: kube-flannel-ds
  namespace: kube-system
  labels:
    tier: node
    app: flannel
spec:
  selector:
    matchLabels:
      app: flannel
  template:
    metadata:
      labels:
        tier: node
        app: flannel
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: kubernetes.io/os
                    operator: In
                    values:
                      - linux
      hostNetwork: true
      priorityClassName: system-node-critical
      tolerations:
        - operator: Exists
          effect: NoSchedule
      serviceAccountName: flannel
      initContainers:
        - name: install-cni
          image: quay.io/coreos/flannel:v0.13.1-rc2
          command:
            - cp
          args:
            - -f
            - /etc/kube-flannel/cni-conf.json
            - /etc/cni/net.d/10-flannel.conflist
          volumeMounts:
            - name: cni
              mountPath: /etc/cni/net.d
            - name: flannel-cfg
              mountPath: /etc/kube-flannel/
      containers:
        - name: kube-flannel
          image: quay.io/coreos/flannel:v0.13.1-rc2
          command:
            - /opt/bin/flanneld
          args:
            - --ip-masq
            - --kube-subnet-mgr
          resources:
            requests:
              cpu: "100m"
              memory: "50Mi"
            limits:
              cpu: "100m"
              memory: "50Mi"
          securityContext:
            privileged: false
            capabilities:
              add: ["NET_ADMIN", "NET_RAW"]
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
          volumeMounts:
            - name: run
              mountPath: /run/flannel
            - name: flannel-cfg
              mountPath: /etc/kube-flannel/
      volumes:
        - name: run
          hostPath:
            path: /run/flannel
        - name: cni
          hostPath:
            path: /etc/cni/net.d
        - name: flannel-cfg
          configMap:
            name: kube-flannel-cfg
```

### kubeadm_join.sh

```sh
#!/bin/bash

sudo kubeadm join ip:6443 --token wxjhuh.z9ru6bbz990m7v0i --discovery-token-ca-cert-hash sha256:8809f1cc27401d704faa104008a48034b51e8e5ef8f4a8f33f0e267db0124a2f
```

### 一些问题

1.k8s: reset 后出现 coredns 一直处于 ContainerCreating 状态

[解决方法 1](https://bbs.huaweicloud.com/forum/thread-60243-1-1.html)

2.k8s: connet to localhost:8080 failed

kubeadm init 成功后，执行`sudo kubectl apply -f kube-flannel.yml`，出现错误：

```sh
The connection to the server localhost:8080 was refused - did you specify the right host or port?
```

[issue](https://github.com/kubernetes/kubernetes/issues/44665)

原来是因为没有执行`export KUBECONFIG=/etc/kubernetes/admin.conf`命令的时候使用 kubectl 是不能在前面加 sudo 的。

所以，直接执行`kubectl apply -f kube-flannel.yml`就正常了。

#### 基于 wsl2 搭建

[基于 wsl2 搭建](https://www.qikqiak.com/post/deploy-k8s-on-win-use-wsl2/)

### Deploy

k8s: 部署应用之 deployment service pod

#### 问题 1

在创建 deployment 时，如果想从本地获取镜像，需要将 yaml 配置文件里的`imagePullPolicy`，镜像拉取机制，从`Always`改为`IfNotPresent`。

#### 问题 2

怎么让 pod 里的应用访问到非集群内部的本机数据库实例呢？

通过新建一个 service，指定为`type:ExternalName`或`Endpoints`，并把想要访问的数据库实例信息写到配置上。

[参照 1](https://juejin.cn/post/6844903693918306312)

[参照 2](https://www.kubernetes.org.cn/4317.html)

[参照 3](https://www.cnblogs.com/ericnie/p/7560280.html)

[参照 4](https://juejin.cn/post/6844903693918306312)

[参照 5](https://www.cnblogs.com/kuku0223/p/10898068.html)

解决：在本机的数据库貌似不能访问到，只好在另外的虚拟机上安装 db，然后配置该虚拟机的 ip 地址到 endpoint。

#### 使用 service 将应用暴露到公网

```sh
# 针对deployment hello-world以NodePort方式、example-service名称、端口是port-value暴露应用到公网
kubectl expose deployment hello-world --type=NodePort --name=example-service -- port=port-value
```

通过公网访问时，需要先使用`kubectl describe svc`拿到 example-service 服务的 NodePort 值，再结合机器自身的 ip 地址，组成 ip:NodePort 值访问，如：<http://192.168.9.16:32256/>。

```sh
jd@wnode1:~/Project/jdnote$ kubectl describe svc
Name:                     jdnote-server-service
Namespace:                default
Labels:                   app=jdnote-server
Annotations:              <none>
Selector:                 app=jdnote-server
Type:                     NodePort
IP Families:              <none>
IP:                       192.168.1.198
IPs:                      192.168.1.198
Port:                     <unset>  8890/TCP
TargetPort:               8890/TCP
NodePort:                 <unset>  32256/TCP # 这个值作为端口
Endpoints:                10.46.129.11:8890
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

```sh
jd@wnode1:~/Project/jdnote$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:cf:aa:13 brd ff:ff:ff:ff:ff:ff
    inet 192.168.9.16/24 brd 192.168.9.255 scope global dynamic enp0s3 # 192.168.9.16作为ip
       valid_lft 9843sec preferred_lft 9843sec
    inet6 fe80::a00:27ff:fecf:aa13/64 scope link
       valid_lft forever preferred_lft forever
```

[参照](https://kubernetes.io/zh/docs/tasks/access-application-cluster/service-access-application-cluster/)

[网络介绍](https://dominik-tornow.medium.com/kubernetes-networking-22ea81af44d0)

#### k8s: ingress

[k8s: ingress](https://kubernetes.io/zh/docs/concepts/services-networking/ingress/)

让外网能访问到 k8s 里的应用。

单个服务：

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress # kind必须指定为ingress
metadata:
  name: minimal-ingress # 名称
  annotations: # 注解
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules: # 规则
    - http:
        paths: # 路径
          - path: /testpath
            pathType: Prefix
            backend: # 重定向到service
              service:
                name: test # 必须存在test service
                port:
                  number: 80 # service的端口
```

多个服务：

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-wildcard-host
spec:
  rules:
    - host: "foo.bar.com" # 域名，表示该域名的流量将被转到配置的service
      http:
        paths:
          - pathType: Prefix
            path: "/bar"
            backend:
              service:
                name: service1
                port:
                  number: 80
    - host: "*.foo.com"
      http:
        paths:
          - pathType: Prefix
            path: "/foo"
            backend:
              service:
                name: service2
                port:
                  number: 80
```

查看：`kubectl describe ingress minimal-ingress`

[可选 ingress 控制器](https://kubernetes.io/zh/docs/concepts/services-networking/ingress-controllers/)

Ingress 只是 Kubernetes 中的一种配置信息；Ingress Controller 才是监听 80/443 端口，并根据 Ingress 上配置的路由信息执行 HTTP 路由转发的组件。

[traefik 提供的 k8s 控制器](https://doc.traefik.io/traefik/providers/kubernetes-ingress/)

[使用 traefik](https://jimmysong.io/kubernetes-handbook/practice/traefik-ingress-installation.html)

## list-watch 模式

[来源](https://zhuanlan.zhihu.com/p/59660536)

> 谈谈 List-Watch 的设计理念
>
> 当设计优秀的一个异步消息的系统时，对消息机制有至少如下四点要求：
>
> > 消息可靠性
> >
> > 消息实时性
> >
> > 消息顺序性
> >
> > 高性能
> >
> > 首先消息必须是可靠的，list 和 watch 一起保证了消息的可靠性，避免因消息丢失而造成状态不一致场景。
> >
> > 具体而言，list API 可以查询当前的资源及其对应的状态(即期望的状态)，客户端通过拿期望的状态和实际的状态进行对比，纠正状态不一致的资源。Watch API 和 apiserver 保持一个长链接，接收资源的状态变更事件并做相应处理。如果仅调用 watch API，若某个时间点连接中断，就有可能导致消息丢失，所以需要通过 list API 解决消息丢失的问题。从另一个角度出发，我们可以认为 list API 获取全量数据，watch API 获取增量数据。虽然仅仅通过轮询 list API，也能达到同步资源状态的效果，但是存在开销大，实时性不足的问题。
>
> 消息必须是实时的，list-watch 机制下，每当 apiserver 的资源产生状态变更事件，都会将事件及时的推送给客户端，从而保证了消息的实时性。
>
> 消息的顺序性也是非常重要的，在并发的场景下，客户端在短时间内可能会收到同一个资源的多个事件，对于关注最终一致性的 K8S 来说，它需要知道哪个是最近发生的事件，并保证资源的最终状态如同最近事件所表述的状态一样。K8S 在每个资源的事件中都带一个 resourceVersion 的标签，这个标签是递增的数字，所以当客户端并发处理同一个资源的事件时，它就可以对比 resourceVersion 来保证最终的状态和最新的事件所期望的状态保持一致。
>
> List-watch 还具有高性能的特点，虽然仅通过周期性调用 list API 也能达到资源最终一致性的效果，但是周期性频繁的轮询大大的增大了开销，增加 apiserver 的压力。而 watch 作为异步消息通知机制，复用一条长链接，保证实时性的同时也保证了性能。

## 调试

golang: dlv 调试 k8s 容器里的 go 进程

### 使用容器 exec 进行调试

如果 容器镜像 包含调试程序(dlv, gdb)， 比如从 Linux 和 Windows 操作系统基础镜像构建的镜像，你可以使用 kubectl exec 命令 在特定的容器中运行一些命令：

`kubectl exec ${POD_NAME} -c ${CONTAINER_NAME} -- ${CMD} ${ARG1} ${ARG2} ... ${ARGN}`

> 说明： -c ${CONTAINER_NAME} 是可选择的。如果 Pod 中仅包含一个容器，就可以忽略它。

例如，要查看正在运行的 Cassandra pod 中的日志，可以运行：

`kubectl exec cassandra -- cat /var/log/cassandra/system.log`

你可以在 kubectl exec 命令后面加上 -i 和 -t 来运行一个连接到你的终端的 Shell，比如：

`kubectl exec -it cassandra -- sh`

[参照 1](https://kubernetes.io/zh/docs/tasks/debug-application-cluster/debug-running-pod/)

[参照 2](https://kubernetes.io/zh/docs/tasks/debug-application-cluster/get-shell-running-container/)

### 尝试

```sh
$ kubectl exec -it skm-7fb6bd989c-kg9xj -n skm-system -- sh # 使用 kubectl 进入容器的 sh 界面

sh-4.4# dlv exec /root/Project/rancher/bin/rancher # exec 报错：版本太旧
Version of Delve is too old for this version of Go (maximum supported version 1.13, suppress this error with --check-go-version=false)
sh-4.4# dlv attach /root/Project/rancher/bin/rancher # attach 需要指定 pid
Invalid pid: /root/Project/rancher/bin/rancher
sh-4.4# dlv attach 653                              # 通过 ps aux 拿到 pid
Type 'help' for list of commands.
(dlv)

```

[dlv 使用](https://chai2010.cn/advanced-go-programming-book/ch3-asm/ch3-09-debug.html)
