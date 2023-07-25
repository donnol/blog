---
author: "jdlau"
date: 2023-07-25
linktitle: Windows restart remote service
menu:
next:
prev:
title: Windows restart remote service
weight: 10
categories: ['Windows']
tags: ['restart-remote-service']
---

Windows如何方便的重启远程服务器里的服务，在不使用远程连接的情况下？

注意：服务(名称：service-name)已在远程机器上创建。

1. 在本机新建`映射网络驱动器`.

2. 打开映射好的文件夹，在其中添加`bat`和`ps1`文件:

`restart.bat`:

```bat
@echo off
for %%i in (service-name) do (
    echo the service '%%i' is being starting...
    sc query %%i
    net stop %%i
    net start %%i
    sc query %%i
    echo service '%%i' started.
)
pause
```

`restart.ps1`:

```sh
$Username = 'Name'
$Password = 'Password'
$pass = ConvertTo-SecureString -AsPlainText $Password -Force
$Cred = New-Object System.Management.Automation.PSCredential -ArgumentList $Username,$pass
Invoke-Command -ComputerName [Remote-IP] -ScriptBlock { Get-Service WinRM } -credential $Cred
Invoke-Command -ComputerName [Remote-IP] -credential $Cred -ScriptBlock { E:\win-dms4\restart.bat } # 此处指定上述bat文件在远程机器上的决对路径
Read-Host -Prompt "Press Enter to exit"
```

完成后，只要在本机打开映射的文件夹，使用`PowerShell`执行`restart.ps1`脚本，即可快速重启指定服务。
