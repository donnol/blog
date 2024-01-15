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
