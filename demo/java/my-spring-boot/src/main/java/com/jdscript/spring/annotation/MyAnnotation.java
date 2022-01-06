package com.jdscript.spring.annotation;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;

// 定义注解
@Retention(RetentionPolicy.RUNTIME) // 没有指定'RUNTIME'的话，反射会拿不到
public @interface MyAnnotation {
    String value() default "this is an annotation by me";
    Class<? extends Object> classValue();
}
