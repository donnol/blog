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

## AOP

面向切面编程（AOP: Aspect Oriented Program）。

### 划分，重复，复用

我们知道，面向对象的特点是继承、多态和封装。而封装就要求将功能**分散**到不同的对象中去，这在软件设计中往往称为职责分配。实际上也就是说，让不同的类设计不同的方法。这样代码就分散到一个个的类中去了。这样做的好处是降低了代码的复杂程度，使类可重用。

> 出现的问题：
>
> 但是人们也发现，在分散代码的同时，也增加了代码的重复性。什么意思呢？比如说，我们在两个类中，可能都需要在每个方法中做日志。按面向对象的设计方法，我们就必须在两个类的方法中都加入日志的内容。也许他们是完全相同的，但就是因为面向对象的设计让类与类之间无法联系，而不能将这些重复的代码统一起来。

想法1：

> 也许有人会说，那好办啊，我们可以将这段代码写在一个独立的类独立的方法里，然后再在这两个类中调用。但是，这样一来，这两个类跟我们上面提到的独立的类就有耦合了，它的改变会影响这两个类。

那么，有没有什么办法，能让我们在需要的时候，随意地加入代码呢？

> 这种在运行时，动态地将代码切入到类的指定方法、指定位置上的编程思想就是面向切面的编程。 
>
> 一般而言，我们管切入到指定类指定方法的代码片段称为切面，而切入到哪些类、哪些方法则叫切入点。

有了AOP，我们就可以把几个类共有的代码，抽取到一个切片中，等到需要时再切入对象中去，从而改变其原有的行为。

OOP从**横向**上区分出一个个的类来，而AOP则从**纵向**上向对象中加入特定的代码。

从技术上来说，AOP基本上是通过**代理**机制实现的。

![AOP](/image/AOP.png)

## Go实现AOP -- 层间代理

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

```go
type userImpl struct {
        userStore IUserStore

        // 增加其它store
        roleStore IRoleStore
}

func (impl userImpl) CheckUser(ctx context.Context, id int) error {
        begin := time.Now()
        user, err := impl.userStore.GetByID(ctx, id)
        if err != nil {
                return err
        }
        fmt.Println(time.Since(begin)) // 统计耗时

        // 使用user数据做一些操作
        _ = user

        // 获取角色具体信息
        {
                begin := time.Now()
                role, err := impl.roleStore.GetByID(ctx, user.RoleId)
                if err != nil {
                        return err
                }
                _ = role
                fmt.Println(time.Since(begin)) // 统计耗时
        }

        // 可能会有更多`Store`
}
```

可以看到，当我们新增了`roleStore`之后，如果要分别统计不同`Store`的方法调用的耗时，将会非常麻烦。这时有人会说，那为什么不把耗时统计放到`Store`的方法实现里呢？或者使用一个方法来封装耗时统计：

```go
func WrapUsedTime[R any](f func() (R, error)) (R, error) {
        begin := time.Now()
        r, err := f()
        if err != nil {
                return r, err
        }
        fmt.Println(time.Since(begin)) // 统计耗时

        return r, nil
}
```

这样做，当然可以。但是，依然很繁琐，特别是在业务很复杂，调用的方法很多的时候。

更重要的一点是，我们应该专注于业务逻辑的开发和测试，通用的东西应该交由框架来实现。这也是`AOP`(面向切面)思想的一个很重要的观点。

其实，在接口开发时用的中间件，也是一种`AOP`实现，但是，中间件的函数签名是固定的，参数类型、参数个数、结果类型和结果个数都是需要事先确定的。但实际中的方法是各种各样的，类型和数量都不尽相同。所以，我们现在要做的是一个通用的`AOP`代理。

```go
// 比如，http包的HandlerFunc，它的签名就是这样的，两个参数，参数类型分别如下：ResponseWriter, *Request
type HandlerFunc func(ResponseWriter, *Request)

// 而我们面临的是：
func GetById(id uint64) (User, error)
func ListByTime(begin, end time.Time) ([]User, error)
```

这时，如果有一个代理能帮我们拦截`store`的方法调用，在调用前后添加上耗时统计，势必能大大提升我们的工作效率。

比如：

```go
// 将函数抽象为func(args []interface{}) []interface{}，
// 用[]interface{}来装所有的参数和结果
func Around(f func(args []interface{}) []interface{}, args []interface{}) []interface{} {
        begin := time.Now()
        r := f(args)
        fmt.Println(time.Since(begin)) // 统计耗时

        return r
}
```

这只是一个简单的包装函数，怎么能将它与上面的接口联系到一起呢？

## [接口，Mock，Around](https://github.com/donnol/tools/blob/master/inject/proxy.go)

```go

func (impl *proxyImpl) around(provider any, mock any, arounder Arounder) any {
	if mock == nil {
		return provider
	}

	mockValue := reflect.ValueOf(mock)
	mockType := mockValue.Type()
	if mockType.Kind() != reflect.Ptr && mockType.Elem().Kind() != reflect.Struct {
		return provider
	}

	// provider有参数，有返回值
	pv := reflect.ValueOf(provider)
	pvt := pv.Type()
	if pvt.Kind() != reflect.Func {
		panic("provider不是函数")
	}

	// 使用新的类型一样的函数
	// 在注入的时候会被调用
	return reflect.MakeFunc(pvt, func(args []reflect.Value) []reflect.Value {

		result := pv.Call(args)

		if len(result) == 0 {
			return result
		}

		firstOut := result[0]
		firstOutType := firstOut.Type()

		if !mockType.Implements(firstOutType) {
			panic(fmt.Errorf("mock not Implements interface"))
		}

		// 根据返回值的类型(mock)生成新的类型，其中新类型的方法均加上钩子
		// 注意：生成的不是接口，是实现了接口的类型
		if firstOutType.Kind() == reflect.Interface {

			newValue := reflect.New(mockType.Elem()).Elem()
			newValueType := newValue.Type()

			// field设置
			for i := 0; i < newValueType.NumField(); i++ {
				field := newValue.Field(i)
				fieldType := newValueType.Field(i)

				var name = fieldType.Name
				for _, suffix := range MockFieldNameSuffixes {
					name = strings.TrimSuffix(name, suffix)
				}

				method := firstOut.MethodByName(name)
				methodType, ok := firstOutType.MethodByName(name)
				if !ok {
					methodTag, ok := fieldType.Tag.Lookup("method")
					if !ok {
						panic(fmt.Errorf("找不到名称对应的方法"))
					}
					debug.Printf("tag: %+v\n", methodTag)
					name = methodTag

					method = firstOut.MethodByName(name)
					methodType, ok = firstOutType.MethodByName(name)
					if !ok {
						panic(fmt.Errorf("使用tag也找不到名称对应的方法"))
					}
				}
				debug.Printf("method: %+v\n", method)

				pctx := ProxyContext{
					PkgPath:       firstOutType.PkgPath(),
					InterfaceName: firstOutType.Name(),
					MethodName:    methodType.Name,
				}
				debug.Printf("pctx: %+v\n", pctx)

				// newMethod会在实际请求时被调用
				// 当被调用时，newMethod内部就会调用绑定好的Arounder，然后将原函数method和参数args传入
				// 在Around方法执行完后即可获得结果
				newMethod := reflect.MakeFunc(methodType.Type, func(args []reflect.Value) []reflect.Value {
					var result []reflect.Value

					debug.Printf("args: %+v\n", args)

					// Around是对整个结构的统一包装，如果需要对不同方法做不同处理，可以根据pctx里的方法名在Around接口的实现里做处理
					result = arounder.Around(pctx, method, args)

					debug.Printf("result: %+v\n", result)

					return result
				})

				field.Set(newMethod)
			}

			result[0] = newValue.Addr().Convert(firstOutType)
		}

		return result
	}).Interface()
}
```

可以看到，主要的方法是`around(provider interface{}, mock interface{}, arounder Arounder) interface{}`，
其中`provider`参数是类似`NewXXX() IXXX`的函数，而`mock`是`IXXX接口`的一个实现，最后的`Arounder`是
拥有方法`Around(pctx ProxyContext, method reflect.Value, args []reflect.Value) []reflect.Value`的接口。

## [这里的示例](https://github.com/donnol/tools/blob/master/inject/proxy_test.go)

可以看到，mock结构是长这样的：

```go
type UserSrvMock struct {
	CheckUserFunc func(ctx context.Context, id int) error
}
```

所以，为了提升开发效率，我还写了一个[工具](https://github.com/donnol/tools)，用来根据接口生成相应的`mock`结构体。

> 安装：`go install github.com/donnol/tools/cmd/tbc@latest`.
>
> 使用：`tbc mock -p=github.com/dominikbraun/graph --mode=offsite`.
>
> 上述命令会解析`graph`包，获取包里的公开接口，然后生成对应的`Mock`结构体，生成的代码保存在当前目录的`mock.go`文件里。

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
		PkgPath:       "接口所在包路径，如：github.com/donnol/tools/inject",
		InterfaceName: "接口名，如：IUserSrv",
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

最后，一个支持任意函数类型的、既能添加通用逻辑，又能添加定制逻辑的`proxy`就完成了。

## 对于任意函数调用通过替换ast节点来添加Proxy

`normal.go`:

```go
package proxy

import (
	"log"
)

func A(ctx any, id int, args ...string) (string, error) {
	log.Printf("arg, ctx: %v, id: %v, args: %+v\n", ctx, id, args)
	return "A", nil
}
func C() {
	args := []string{"a", "b", "c", "d"}
	r1, err := A(1, 1, args...)
	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}
	log.Printf("r1: %v\n", r1)
}
```

在上述代码中，`C`函数调用了`A`函数，那么，现在我想在这个调用前后添加耗时统计，该怎么办呢？

```go
// 添加耗时统计
func C() {
        begin := time.Now()

	args := []string{"a", "b", "c", "d"}
	r1, err := A(1, 1, args...)
	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}
	log.Printf("r1: %v\n", r1)

        log.Printf("used time: %v\n", time.Since(begin))
}
```

如果，我能生成一个`AProxy`函数，里面包含有耗时统计等逻辑，再把`C`对`A`的调用改为对`Aproxy`的调用，是不是就非常方便了呢！

```sh
# 执行命令，生成代码
tbc genproxy -p ./parser/testtype/proxy/ --func A
```

`gen_proxy.go`:

```go
package proxy

import (
	"log"
	"time"
)

// 生成A的Proxy
func AProxy(ctx any, id int, args ...string) (string, error) {
	begin := time.Now()

	var r0 string
	var r1 error

	r0, r1 = A(ctx, id, args...)

	log.Printf("used time: %v\n", time.Since(begin))

	return r0, r1
}
```

`normal.go`:

```go
package proxy

import (
	"log"
)

func A(ctx any, id int, args ...string) (string, error) {
	log.Printf("arg, ctx: %v, id: %v, args: %+v\n", ctx, id, args)
	return "A", nil
}
func C() {
	args := []string{"a", "b", "c", "d"}
        // 此处对A的调用就被替换为对AProxy的调用了
	r1, err := AProxy(1, 1, args...)
	if err != nil {
		log.Printf("err: %v\n", err)
		return
	}
	log.Printf("r1: %v\n", r1)
}
```

不过，这种方式会修改用户编写的源代码，使用时请注意。

[代码实现详见](https://github.com/donnol/tools/blob/feat/inject-proxy-caller/cmd/tbc/main.go#L404)
