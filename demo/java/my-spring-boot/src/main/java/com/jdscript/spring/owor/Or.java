package com.jdscript.spring.owor;

public class Or extends Owor implements O {
    @Override // 编译器会校验写的方法在父类中是否存在
    public String name() {
        super.name();
        return "or";
    }

    public String name(int a) {
        return "overload";
    }

    // @Override // 如果添加Override，会报错： The method a() of type Or must override or implement a supertype method
    public String a() {
        return "a";
    }

    public String o() {
        return "o";
    }
}
