# wsl2的一些配置

以下是在wsl虚拟机里面的配置

## 配置

```sh
$ cat /etc/wsl.conf 
[boot]
systemd=true # 启用systemd

[interop]
appendWindowsPath = false

[network]
generateResolvConf = false # 关闭自动生成resolv.conf
```

## 更新resolv.conf

```sh
$ cat Script/generateResolvConf.sh 
#!/bin/sh

echo 'nameserver 192.168.8.44' > /etc/resolv.conf
```

以`systemd service`在开机时执行脚本：

```sh
$ cat /etc/systemd/system/generateResolvConf.service
[Unit]
Description=Run generateResolvConf.sh to set dns of wsl2 at Startup

[Service]
ExecStart=/home/jd/Script/generateResolvConf.sh

[Install]
WantedBy=default.target
```

## 更新wslhost

wsl启动时，注册虚拟机ip到主机hosts：

```sh
$ cat Script/wsl2host.sh 
#!/bin/bash
HOSTS_FILE_WIN='/mnt/c/Windows/System32/drivers/etc/hosts'
TEMP_DIR_PATH='~/tmp/dns'
TEMP_FILE_PATH=${TEMP_DIR_PATH}'/dns_back'

# inetIp=`ifconfig eth0 | grep -o "inet [0-9]*\.[0-9]*\.[0-9]*\.[0-9]* netmask" | cut -f 2 -d " "` # 获取本机ip
inetIp=`ip a | grep eth0 | grep -o "inet [0-9]*\.[0-9]*\.[0-9]*\.[0-9]*" | cut -f 2 -d " "`

nu=`cat -n ${HOSTS_FILE_WIN} | grep localwsl2 | awk '{print $1}'` # 获取已设置dns行号

dnsIp=`cat ${HOSTS_FILE_WIN} | grep -o "[0-9]*\.[0-9]*\.[0-9]*\.[0-9]* localwsl2 #" | cut -f 1 -d " "` # 获取已设置dns ip 

echo "wsl's ip is: ${inetIp}"
echo "win's dns line number is: ${nu}"
echo "win's dnsIp is: ${dnsIp}"

if [ -z ${inetIp} ]
then
        echo 'inet ip is null, please check the command'
fi

set -e
if [ ${nu} ] # 若已设置
then
        if [ -z ${dnsIp} ] || [ ${inetIp} != ${dnsIp} ] # 已设置dns不正确
        then
                echo reset
                if [ ! -e ${TEMP_FILE_PATH} ]
                then
                        echo "will mkdir ${TEMP_DIR_PATH} and file ${TEMP_FILE_PATH}"
                        mkdir -p ${TEMP_DIR_PATH} && touch ${TEMP_FILE_PATH}
                fi
                cp -f "${HOSTS_FILE_WIN}" "${TEMP_FILE_PATH}"
                sed -i "${nu}d" ${TEMP_FILE_PATH} # 删除对应行
                echo "${inetIp} localwsl2 #wsl2 dns config" >> ${TEMP_FILE_PATH} # 重新设置
                cat ${TEMP_FILE_PATH} > ${HOSTS_FILE_WIN}
                echo set success!!!
        fi
else # 未设置
        echo "will append localwsl2 ip:host to windows hosts"
        echo "${inetIp} localwsl2 #wsl2 dns config" >> ${HOSTS_FILE_WIN} # 直接设置
        echo "finish append localwsl2 ip:host to windows hosts"
fi
```

以`systemd service`在开机时执行`wsl2host.sh`脚本：

```sh
$ cat /etc/systemd/system/wsl2host.service 
[Unit]
Description=Run wsl2host.sh to set dns of wsl2 at Startup

[Service]
ExecStart=/home/jd/Script/wsl2host.sh

[Install]
WantedBy=default.target
```

## .wslconfig

```toml
# Settings apply across all Linux distros running on WSL 2
[wsl2]

# Limits VM memory to use no more than 4 GB, this can be set as whole numbers using GB or MB
# memory=4GB 

# Sets the VM to use two virtual processors
# processors=2

# Specify a custom Linux kernel to use with your installed distros. The default kernel used can be found at https://github.com/microsoft/WSL2-Linux-Kernel
# kernel=C:\\temp\\myCustomKernel

# Sets additional kernel parameters, in this case enabling older Linux base images such as Centos 6
# kernelCommandLine = vsyscall=emulate

# Sets amount of swap storage space to 8GB, default is 25% of available RAM
# swap=8GB

# Sets swapfile path location, default is %USERPROFILE%\AppData\Local\Temp\swap.vhdx
# swapfile=C:\\temp\\wsl-swap.vhdx

# Disable page reporting so WSL retains all allocated memory claimed from Windows and releases none back when free
# pageReporting=false

# Turn off default connection to bind WSL 2 localhost to Windows localhost
# localhostforwarding=true

# Disables nested virtualization
# nestedVirtualization=false

# Turns on output console showing contents of dmesg when opening a WSL 2 distro for debugging
# debugConsole=true

# win11才能用，win10会报错：虚拟机或容器 JSON 文档无效。Error code: Wsl/Service/CreateInstance/CreateVm/ConfigureNetworking/0x8037010d
# networkingMode=bridged
# vmSwitch=my-switch
# ipv6=true

[experimental]
autoMemoryReclaim=gradual # 检测空闲 CPU 使用率后，自动释放缓存的内存。 设置为 gradual 以慢速释放，设置为 dropcache 以立即释放缓存的内存。
```

## wsl2 filesystem performance

```sh
$ dd if=/dev/zero of=/mnt/e/testfile bs=1M count=1000
1000+0 records in
1000+0 records out
1048576000 bytes (1.0 GB, 1000 MiB) copied, 3.88478 s, 270 MB/s
```

[issues](https://github.com/microsoft/WSL/issues/4197)
