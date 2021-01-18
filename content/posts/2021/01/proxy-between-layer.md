---
author: "jdlau"
date: 2021-01-17
linktitle: go实现AOP
menu:
next: 
prev: 
title: go实现AOP
weight: 10
categories: ['go']
tags: ['aop']
---

## go实现AOP

假设有store，从数据库获取数据，其中有方法IUserStore.GetByID，传入id参数，返回用户信息:

```go
type IUserStore interface {
        GetByID(ctx context.Context, id int) (User, error)
}
```

另外有service，刚好有用户id并且需要拿到用户信息，于是依赖了上述IUserStore：

```go
type IUserSrv interface {
        CheckUser(ctx context.Context, id int) error // 获取用户信息，然后检查用户某些属性
} 

type userImpl struct {
        userStore IUserStore
}

func (impl userImpl) CheckUser(ctx context.Context, id int) error {
        user, err := impl.userStore.GetByID(ctx, id)
        if err != nil {
                return err
        }

        // 使用user数据做一些操作
        _ = user
}
```

上面所描述的是一个最简单的情况，如果我们要在userImpl.CheckUser里对impl.userStore.GetByID方法调用添加耗时统计，依然十分简单。

```go
func (impl userImpl) CheckUser(ctx context.Context, id int) error {
        begin := time.Now()
        user, err := impl.userStore.GetByID(ctx, id)
        if err != nil {
                return err
        }
        fmt.Println(time.Since(begin)) // 统计耗时

        // 使用user数据做一些操作
        _ = user
}
```

但是，如果方法里调用的类似impl.userStore.GetByID的方法非常之多，逻辑非常之复杂时，这样一个一个的添加，必然非常麻烦、非常累。

这时，如果有一个层间代理能帮我们拦截store的方法调用，在调用前后添加上耗时统计，势必能大大提升我们的工作效率。

比如：

```go
func Around(f func(args []interface{}) []interface{}, args []interface{}) []interface{} {
        begin := time.Now()
        r := f(args)
        fmt.Println(time.Since(begin)) // 统计耗时

        return r
}
```

这只是一个简单的包装函数，怎么能将它与上面的接口联系到一起呢？

## [有兴趣的话，可以看这里的实现](https://github.com/donnol/tools/blob/master/inject/proxy.go)

可以看到，主要的方法是`Around(provider interface{}, mock interface{}, arounder Arounder) interface{}`，
其中provider参数是类似`NewXXX() IXXX`的函数，而mock是IXXX接口的一个实现，最后的Arounder是
拥有方法`Around(pctx ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value`的接口。



## [这里的示例](https://github.com/donnol/tools/blob/master/inject/proxy_test.go)

可以看到，mock结构是长这样的：

```go
type UserMock struct {
	AddFunc        func(name string) int
	GetHelper      func(id int) string `method:"Get"` // 表示这个字段关联的方法是Get
	GetContextFunc func(ctx context.Context, id int) string
}
```

所以，为了提升开发效率，我还写了一个[工具](https://github.com/donnol/tools)，用来根据接口生成相应的mock结构体。
