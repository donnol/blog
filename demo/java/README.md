# java

## spring boot

> Spring Boot makes it easy to create stand-alone, production-grade Spring based Applications that you can "just run".
> 
> -- Spring Boot让创建刚好能运行的独立的生产级应用变得简单。

### 搭建

安装`jdk`、`maven`、`spring`，均要在`PATH`里配置它们的`bin`目录。

初始化`spring boot`

```java
spring init -d='web,actuator' -n=myspring --package-name=com.jdscript.spring my-spring-boot
```

`spring`命令更多选项请看`spring help`

下面是`spring init`可用选项情况：

```sh
$ spring help init
spring init - Initialize a new project using Spring Initializr (start.spring.io)

usage: spring init [options] [location]

Option                       Description
------                       -----------
-a, --artifact-id <String>   Project coordinates; infer archive name (for
                               example 'test')
-b, --boot-version <String>  Spring Boot version (for example '1.2.0.RELEASE')
--build <String>             Build system to use (for example 'maven' or
                               'gradle') (default: maven)
-d, --dependencies <String>  Comma-separated list of dependency identifiers to
                               include in the generated project
--description <String>       Project description
-f, --force                  Force overwrite of existing files
--format <String>            Format of the generated content (for example
                               'build' for a build file, 'project' for a
                               project archive) (default: project)
-g, --group-id <String>      Project coordinates (for example 'org.test')
-j, --java-version <String>  Language level (for example '1.8')
-l, --language <String>      Programming language  (for example 'java')
--list                       List the capabilities of the service. Use it to
                               discover the dependencies and the types that are
                               available
-n, --name <String>          Project name; infer application name
-p, --packaging <String>     Project packaging (for example 'jar')
--package-name <String>      Package name
-t, --type <String>          Project type. Not normally needed if you use --
                               build and/or --format. Check the capabilities of
                               the service (--list) for more details
--target <String>            URL of the service to use (default: https://start.
                               spring.io)
-v, --version <String>       Project version (for example '0.0.1-SNAPSHOT')
-x, --extract                Extract the project archive. Inferred if a
                               location is specified without an extension

examples:
    To list all the capabilities of the service:
        $ spring init --list

    To creates a default project:
        $ spring init

    To create a web my-app.zip:
        $ spring init -d=web my-app.zip

    To create a web/data-jpa gradle project unpacked:
        $ spring init -d=web,jpa --build=gradle my-dir
```

### 开发

### 构建

### 部署

[官网](https://spring.io/projects/spring-boot)

### 运行

#### 监控

#### 反馈

## android开发入门

1. [下载studio](https://developer.android.com/studio/)

2. 解压

3. 安装：`bash ./bin/studio.sh`；下载依赖时需要设置代理。

4. 新建项目，除了修改项目路径，其余使用默认设置。

studio加载项目的时候因为要下依赖，速度非常之慢，尝试修改配置`build.gradle`：

```gradle
    repositories {
        maven{ url 'https://maven.aliyun.com/repository/central' }
        maven{ url 'https://maven.aliyun.com/repository/public' }
        maven{ url 'https://maven.aliyun.com/repository/gradle-plugin'}
        maven{ url 'https://maven.aliyun.com/repository/grails-core'}
    }
```

在加载模拟器的时候，因为用的是虚拟机，需要开启虚拟机CPU的`VT-x`特性。

在虚拟机配置里开启了该特性之后，又报错：`此平台不支持虚拟化的 Intel VT-x/EPT`，暂时先不管它。
