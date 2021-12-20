---
author: "jdlau"
date: 2021-12-13
linktitle: redis sds
menu:
next:
prev:
title: redis sds
weight: 10
categories: ['redis']
tags: ['sds']
---

## 简单动态字符串

结构：

| len                  | alloc                                                  | flag                | buf                  |
| -------------------- | ------------------------------------------------------ | ------------------- | -------------------- |
| 长度(已使用空间大小) | 分配(总共空间大小：buf 的大小减 1 -- '\0'字符占用了 1) | 标记(sdshdr 的类型) | 真正存储字符串的地方 |

文件：

sds.h, sdsalloc.h, sds.c.

定义：

```c++
typedef char *sds;
```

根据类型获取长度：

```c++
static inline size_t sdslen(const sds s) {
    unsigned char flags = s[-1];
    switch(flags&SDS_TYPE_MASK) {
        case SDS_TYPE_5:
            return SDS_TYPE_5_LEN(flags);
        case SDS_TYPE_8:
            return SDS_HDR(8,s)->len;
        case SDS_TYPE_16:
            return SDS_HDR(16,s)->len;
        case SDS_TYPE_32:
            return SDS_HDR(32,s)->len;
        case SDS_TYPE_64:
            return SDS_HDR(64,s)->len;
    }
    return 0;
}
```

新建：

```c++
/* Create a new sds string with the content specified by the 'init' pointer
 * and 'initlen'.
 * If NULL is used for 'init' the string is initialized with zero bytes.
 * If SDS_NOINIT is used, the buffer is left uninitialized;
 *
 * The string is always null-termined (all the sds strings are, always) so
 * even if you create an sds string with:
 *
 * mystring = sdsnewlen("abc",3);
 *
 * You can print the string with printf() as there is an implicit \0 at the
 * end of the string. However the string is binary safe and can contain
 * \0 characters in the middle, as the length is stored in the sds header. */ // sds的头部存储了它的长度
sds _sdsnewlen(const void *init, size_t initlen, int trymalloc) {
    void *sh; // sds的header
    sds s;
    char type = sdsReqType(initlen); // 根据size大小返回类型
    /* Empty strings are usually created in order to append. Use type 8
     * since type 5 is not good at this. */
    if (type == SDS_TYPE_5 && initlen == 0) type = SDS_TYPE_8;
    int hdrlen = sdsHdrSize(type); // 根据类型返回header长度
    unsigned char *fp; /* flags pointer. */
    size_t usable; // 可用的内存大小

    assert(initlen + hdrlen + 1 > initlen); /* Catch size_t overflow */
    sh = trymalloc?
        s_trymalloc_usable(hdrlen+initlen+1, &usable) :
        s_malloc_usable(hdrlen+initlen+1, &usable); // 三元运算符，根据trymalloc的值选择用哪个alloc函数，分配后，还会标志可用内存大小
    if (sh == NULL) return NULL; // 分配失败返回NULL
    if (init==SDS_NOINIT)
        init = NULL;
    else if (!init)
        memset(sh, 0, hdrlen+initlen+1);
    s = (char*)sh+hdrlen; // 将sh转为char*，赋值给最终字符串s
    fp = ((unsigned char*)s)-1;
    usable = usable-hdrlen-1;
    if (usable > sdsTypeMaxSize(type))
        usable = sdsTypeMaxSize(type);
    switch(type) { // 根据类型决定对s做不同的sds hdr操作
        case SDS_TYPE_5: {
            *fp = type | (initlen << SDS_TYPE_BITS);
            break;
        }
        case SDS_TYPE_8: {
            SDS_HDR_VAR(8,s);
            sh->len = initlen;
            sh->alloc = usable;
            *fp = type;
            break;
        }
        case SDS_TYPE_16: {
            SDS_HDR_VAR(16,s);
            sh->len = initlen;
            sh->alloc = usable;
            *fp = type;
            break;
        }
        case SDS_TYPE_32: {
            SDS_HDR_VAR(32,s);
            sh->len = initlen;
            sh->alloc = usable;
            *fp = type;
            break;
        }
        case SDS_TYPE_64: {
            SDS_HDR_VAR(64,s);
            sh->len = initlen;
            sh->alloc = usable;
            *fp = type;
            break;
        }
    }
    if (initlen && init)
        memcpy(s, init, initlen); // 复制initlen的init内容到s
    s[initlen] = '\0';
    return s;
}
```

[引用 1](https://segmentfault.com/a/1190000025130861#:~:text=SDS%20%28simple,dynamic%20string%29%E6%98%AFRedis%E6%8F%90%E4%BE%9B%E7%9A%84%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%9A%84%E5%B0%81%E8%A3%85%EF%BC%8C%E5%9C%A8redis%E4%B8%AD%E4%B9%9F%E6%98%AF%E5%AD%98%E5%9C%A8%E6%9C%80%E5%B9%BF%E6%B3%9B%E7%9A%84%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%EF%BC%8C%E5%AE%83%E4%B9%9F%E6%98%AF%E5%BE%88%E5%A4%9A%E5%85%B6%E4%BB%96%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E7%9A%84%E5%9F%BA%E7%A1%80%EF%BC%8C%E6%89%80%E4%BB%A5%E6%89%8D%E9%80%89%E6%8B%A9%E5%85%88%E4%BB%8B%E7%BB%8DSDS%E3%80%82)
