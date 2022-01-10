---
author: "jdlau"
date: 2022-01-10
linktitle: 缓存和数据库如何保持一致
menu:
next:
prev:
title: 缓存和数据库如何保持一致
weight: 10
categories: ['cache', 'database']
tags: ['redis', 'mysql']
---

## 缓存读

从缓存读，如果读到了，直接返回；如果读不到，继续去数据库读（singleflight），读到后，更新缓存，返回结果。

## 缓存写

为什么是删缓存，而不是更新缓存呢？

主要是怕两个并发的写操作导致脏数据。

删除缓存和更新磁盘谁先谁后呢？

1.如果先删除缓存，再更新磁盘时的问题：

数据发生了变更，先删除了缓存，然后要去修改数据库，此时还没修改。 -- 删了缓存，未完成数据库修改
另一个请求过来，去读缓存，发现缓存空了，去查询数据库，查到了修改前的旧数据，放到了缓存中。 -- 因为上面的请求里修改数据库的部分还未完成
随后数据变更的程序完成了数据库的修改。-- 这时才完成，可缓存已经填充了之前的旧值了
来到这，数据库和缓存中的数据就不一样了。

2.先更新磁盘，再删除缓存的问题：

先更新数据库，再删除缓存，如果数据库更新了，但是缓存删除失败了，那么缓存中的数据还是旧数据，出现数据不一致

先删除缓存，再更新数据库。如果数据库更新失败了，那么数据库中是旧数据，缓存中是空的，那么数据不会不一致。

比如，一个是读操作，但是没有命中缓存，然后就到数据库中取数据，此时来了一个写操作，写完数据库后，让缓存失效，然后，之前的那个读操作再把老的数据放进去，所以，会造成脏数据。

但，这个case理论上会出现，不过，实际上出现的概率可能非常低，因为这个条件需要发生在读缓存时缓存失效，而且并发着有一个写操作。而实际上数据库的写操作会比读操作慢得多，而且还要锁表，而读操作必需在写操作前进入数据库操作，而又要晚于写操作更新缓存，所有的这些条件都具备的概率基本并不大。

所以，这也就是Quora上的那个答案里说的，要么通过2PC或是Paxos协议保证一致性，要么就是拼命的降低并发时脏数据的概率，而Facebook使用了这个降低概率的玩法，因为2PC太慢，而Paxos太复杂。当然，最好还是为缓存设置上过期时间。

[参照](https://coolshell.cn/articles/17416.html)

## 代码

```go
package cache

// 缓存，一般先将数据从磁盘读出来写到内存里，供用户高速访问，减少读磁盘 -- 快取
// 另有缓冲，将数据先写到内存里，待装满后一次性写入磁盘，可以少写很多次 -- 缓冲
//
// 不难看出，无论快取还是缓冲，都涉及到内存和磁盘的读写。
//
// 首先，对于缓存，目前使用较多的中间件是redis、memcached等，当然也有自己在程序中内置map充当缓存的。
// 那么，下面来看下如何在内存和磁盘之间同步数据：

type Cache interface {
	// exp表示当前时间后的exp秒后过期，传0则无过期
	Set(key, value string, exp int) error
	Del(key string) error
	Get(key string) (value string, err error)
}

type Store interface {
	Create(key, value string, exp int) error
	Delete(key string) error
	Update(key, value string, exp int) error
	Get(key string) (vlaue string, err error)
}

// 用户 有增删改查四个操作，在操作时，对应的缓存和磁盘如何变化呢？
type Client struct {
	cache Cache
	store Store
}

func NewClient(
	cache Cache,
	store Store,
) *Client {
	return &Client{
		cache: cache,
		store: store,
	}
}

const (
	defExp = 300
)

// Add 先写磁盘还是缓存呢？
func (client *Client) Add(key, value string) {
	// 会不会已经在Store里存在了呢？
	// 先从Store Get一次？
	// 一般来说，key都是唯一的：
	// 此时，必须请求一次Store，确认数据不存在；如果此时数据存在，直接返回错误
	if v, err := client.store.Get(key); err == nil && v != "" {
		return
	}

	// 缓存里会不会也有呢？
	// 先从Cache里读一次？
	// 如果能通过业务检查，正常来说，缓存里是没有的；

	// 写磁盘
	client.store.Create(key, value, defExp)

	// 1. 写缓存
	// client.cache.Set(key, value, defExp)
	// 2. 不写，等获取时从磁盘取
}

func (client *Client) Mod(key, value string) {
	// 更新磁盘
	client.store.Update(key, value, defExp)

	// 1. 更新缓存
	// client.cache.Set(key, value, defExp)
	// 2. 不更新，删掉缓存，等获取时从磁盘取 -- Cache Aside Pattern(旁路缓存方案): 一个 lazy 计算的思想，不要每次都重新做复杂的计算，而是让它到需要被使用的时候再重新计算
	client.cache.Del(key)
}

func (client *Client) Del(key string) {
	// 从磁盘删除
	client.store.Delete(key)

	// 1. 从缓存删除
	client.cache.Del(key)
}

func (client *Client) Get(key string) (string, error) {
	// 从Cache获取
	v, err := client.cache.Get(key)
	if err != nil {
		return "", err
	}
	if v == "" {
		// 以SingleFlight方式从Store读:
		// https://pkg.go.dev/golang.org/x/sync/singleflight
		// 这个库的主要作用就是将一组相同的请求合并成一个请求，实际上只会去请求一次，然后对所有的请求返回相同的结果。
		// 在一个请求的时间周期内实际上只会向底层的数据库发起一次请求大大减少对数据库的压力。
		{
			valueFromStore, err := client.store.Get(key)
			if err != nil {
				return "", err
			}
			v = valueFromStore
		}

		// 设置Cache
		client.cache.Set(key, v, defExp)
	}

	return v, nil
}
```
