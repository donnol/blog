---
author: "jdlau"
date: 2021-12-02
linktitle: github action deploy hugo blog
menu:
next: 
prev: 
title: github action deploy hugo blog
weight: 10
categories: ['blog']
tags: ['deploy']
---

## why

为了将视线保持在文章上，减少构建和发布的时间占用。

## what

`github action`是`GitHub`推出的持续集成/持续部署工具，只需要在项目中添加`workflow.yml`配置文件，在其中配置好任务、工作、步骤等，即可在指定动作发生时自动触发编排好的动作。换言之，如果我们在我们的博客仓库里配置了自动将内容打包和发布的`workflow.yml`，那我们就可以把精力集中在文章的编写，然后轻轻地提交推送，即可完成博客地打包和发布，`very easy and smooth`。

## how

在github准备一个blog仓库，用于存放原始信息；再准备一个`github page`仓库，用于存放打包数据。

其中`github page`仓库已开启page，可以通过`github page`设置的域名访问。

[我的blog仓库](https://github.com/donnol/blog)

[我的github page仓库](https://github.com/donnol/donnol.github.io)

### workflow

[这是我结合网络各位英豪所总结出来的一个workflow.yml配置文件](https://github.com/donnol/blog/blob/main/.github/workflows/workflow.yml)

```yaml
name: blog # 做什么都好，别忘了先起个平凡（kuxuan）的名字

on: # 指定触发动作
  push: # 动作是：git push
    branches:
      - main # 指定分支： main

jobs:
  build-deploy:
    runs-on: ubuntu-latest # 基于ubuntu
    steps:
    - uses: actions/checkout@v2 # 切换分支：git checkout
      with:
        submodules: recursive  # Fetch Hugo themes (true OR recursive)
        fetch-depth: 0    # Fetch all history for .GitInfo and .Lastmod

    - name: Setup Hugo # 博客所用的打包和部署工具
      uses: peaceiris/actions-hugo@v2
      with:
        hugo-version: latest

    - name: Build # 打包
      run: hugo --minify --baseURL=https://donnol.github.io # 指定base url，确保构建出来的内容里的超链接都在它里面

    - name: Deploy # 部署
      uses: peaceiris/actions-gh-pages@v3
      with:
        deploy_key: ${{ secrets.ACTIONS_DEPLOY_KEY }} # 这个key非常关键，一言两语很难讲清楚
        external_repository: donnol/donnol.github.io # 我的github page所在的仓库
        PUBLISH_BRANCH: main
        PUBLISH_DIR: ./public # 将本仓库的public目录下的内容提交到github page仓库
        commit_message: ${{ github.event.head_commit.message }} # 提交信息
```

以铜为镜，可以正衣冠；以人为镜，可以明得失；[以史为镜，可以知兴替](https://github.com/donnol/blog/actions)。

### deploy key

1. 使用`ssh-keygen`生产一对非对称秘钥（包含有公钥、私钥）

2. 在`github page`(我这里是`donnol/donnol.github.io`)的仓库的setting里的deploy里添加公钥

3. 在blog仓库setting的`secrets`里添加私钥，注意命名必须是workflow里使用的名称(如上述：ACTIONS_DEPLOY_KEY)

### Q&A

遇到问题不要惊慌，阿Q怕的是强者，如果你示弱，结果可想而知。

当然，实在搞不懂，也可以在issue里提问，本人不负责任地想回就回。

## 温馨提示

如果想知道更详细的信息，请自行搜索关键词，网络大神比比皆是，学习资料处处有售，生活实践时时待你。
