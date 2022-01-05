package com.jdscript.spring.annotation;

// 定义注解
public @interface MyAnnotation {
    String value() default "this is an annotation by me";
    Class<? extends Object> classValue();
}
