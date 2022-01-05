package com.jdscript.spring;

import javax.annotation.PostConstruct;

import com.jdscript.spring.annotation.MyAnnotation;
import com.jdscript.spring.owor.Or;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MyAnnotation(value="main", classValue=MyspringApplication.class) // 使用注解
public class MyspringApplication {
	@Value("${my.user.name}") // ${user.name} 获取的是服务器的当前登录用户名称
	private String name;

	private void printName() {
		System.out.println("===== name: "+this.name);
	}

	@PostConstruct // executed after dependency injection is done -- 依赖注入后执行
	public void init (){
 		this.printName() ;
	}

	// TODO:获取注解，并根据注解内容做出特定动作

	public static void main(String[] args) {
		Or or = new Or();
		System.out.println(or);

		SpringApplication.run(MyspringApplication.class, args);
	}

}
