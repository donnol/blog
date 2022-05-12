---
author: "jdlau"
date: 2021-01-17
linktitle: Go实现AOP
menu:
next: 
prev: 
title: Go实现AOP
weight: 10
categories: ['go']
tags: ['aop', 'proxy']
---

## Go实现AOP

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
type UserSrvMock struct {
	CheckUserFunc func(ctx context.Context, id int) error
}
```

所以，为了提升开发效率，我还写了一个[工具](https://github.com/donnol/tools)，用来根据接口生成相应的mock结构体。

## 代码生成替代反射

在上面描述的`Around`实现里，依赖了`reflect`包里的`reflect.Value.Call`方法：

```go
func (v Value) Call(in []Value) []Value
```

而这个方法的性能是比直接方法调用差的，因此，能不能用代码生成来替代它呢？

再回过头来看一下，我们通过`provider`新建一个对象，这个对象带有我们需要使用的方法：

```go
func NewIUserSrv(userStore IUserStore) IUserSrv {
        return &userImpl{
                userStore: userStore,
        }
}
```

如果我们把`provider`改为：

```go
func NewIUserSrv(userStore IUserStore, withProxy bool) IUserSrv {
        base := &userImpl{
                userStore: userStore,
        }
        if withProxy { // 控制是否使用proxy
                return getIUserSrvProxy(base)
        }
        return base
}

func getIUserSrvProxy(base IUserSrv) *UserSrvMock {
        return &UserSrvMock{
               CheckUserFunc: func(ctx context.Context, id int) error {
                       var r0 error

                       // 这里不就可以添加逻辑了吗

                       r0 = base.CheckUser(ctx, id)

                       // 这里不就可以添加逻辑了吗

                       return r0
               },
        }
}
```

这样，不就可以在调用该方法前后添加逻辑了吗？

如果接口的方法很多，并且添加的逻辑都一样，我们就需要考虑使用代码生成来提高开发效率了：

```go
// 生成getIUserSrvProxy函数
func getIUserSrvProxy(base IUserSrv) *UserSrvMock {
        return &UserSrvMock{
                CheckUserFunc: func(ctx context.Context, id int) error {
                        // 通用逻辑：耗时统计
                        _gen_begin := time.Now()

                        var _gen_r0 error

                        _gen_ctx := UserSrvMockCheckUserProxyContext // 生成Mock时一并生成
                        _gen_cf, _gen_ok := _gen_customCtxMap[_gen_ctx.Uniq()] // _gen_customCtxMap：全局map，存储用户自定义proxy
                        if _gen_ok {
                                // 收集参数
                                _gen_params := []any{}

                                _gen_params = append(_gen_params, ctx)

                                _gen_params = append(_gen_params, id)

                                _gen_res := _gen_cf(_gen_ctx, base.CheckUser, _gen_params)

                                // 结果断言
                                _gen_tmpr0, _gen_exist := _gen_res[0].(error)
                                if _gen_exist {
                                        _gen_r0 = _gen_tmpr0
                                }

                        } else {
                                // 原始调用
                                _gen_r0 = base.CheckUser(ctx, id)
                        }

                        log.Printf("[ctx: %s]used time: %v\n", _gen_ctx.Uniq(), time.Since(_gen_begin))

                        return _gen_r0
		},
        }
}

var (
        userSrvMockCommonProxyContext = inject.ProxyContext{
		PkgPath:       "接口所在包",
		InterfaceName: "包名",
	}
	UserSrvMockCheckUserProxyContext = func() (pctx inject.ProxyContext) {
		pctx = userSrvMockCommonProxyContext
		pctx.MethodName = "CheckUser" // 方法名
		return
	}()
)

var (
	_gen_customCtxMap = make(map[string]inject.CtxFunc)
)

// 通过调用这个方法注册自定义proxy函数
func RegisterProxyMethod(pctx inject.ProxyContext, cf inject.CtxFunc) {
	_gen_customCtxMap[pctx.Uniq()] = cf
}

func main() {
        RegisterProxyMethod(UserSrvMockCheckUserProxyContext, func(ctx ProxyContext, method any, args []any) (res []any) {
		log.Printf("custom call")

                // 从any断言回具体的函数、参数
		f := method.(func(ctx context.Context, id int) error)
		a0 := args[0].(context.Context)
                a1 := args[1].(id)

                // 调用
		r1 := f(a0, a1)
		res = append(res, r1)

		return res
	})
}
```

最后，一个既能添加通用逻辑，又能添加定制逻辑的`proxy`就完成了。
