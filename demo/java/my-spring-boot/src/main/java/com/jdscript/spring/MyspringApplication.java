package com.jdscript.spring;

import java.lang.annotation.Annotation;
import java.lang.reflect.Field;

import javax.annotation.PostConstruct;

import com.jdscript.spring.annotation.MyAnnotation;
import com.jdscript.spring.owor.Or;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MyAnnotation(value = "main", classValue = MyspringApplication.class) // 使用注解
public class MyspringApplication {
	@Value("${my.user.name}") // ${user.name} 获取的是服务器的当前登录用户名称
	private String name;

	private void printName() {
		System.out.println("===== name: " + this.name);
	}

	@PostConstruct // executed after dependency injection is done -- 依赖注入后执行
	public void init() {
		this.printName();
		this.getAnnotation();
	}

	// 获取注解，并根据注解内容做出特定动作
	public void getAnnotation() {
		Class<?> classz = MyspringApplication.class;
		try {
			Annotation[] annos = classz.getAnnotations();
			System.out.println(annos.length); // 为啥这里是1呢？不是有两个注解吗？
			for (Annotation annotation : annos) {
				System.out.println(annotation);
			}
			
			Field field = classz.getField("name");
			field.getAnnotations();
		} catch (Exception e) {

		} finally {

		}
	}

	public static void main(String[] args) {
		Or or = new Or();
		System.out.println(or);

		SpringApplication.run(MyspringApplication.class, args);
	}

}
