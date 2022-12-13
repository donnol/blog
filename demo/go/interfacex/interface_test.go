package interfacex_test

import (
	"fmt"
	"go/ast"
	"testing"

	"github.com/donnol/blog/demo/go/interfacex"
)

type M struct {
	interfacex.I
}

type m struct {
}

func (m m) Method() {
	fmt.Println("call method from m")
}

func TestI(t *testing.T) {
	m := M{
		I: m{},
	}
	m.Method()
}

type MU struct {
	interfacex.IunexportMethod
}

type mu struct {
}

func (mu mu) Method() {
	fmt.Println("call method from mu")
}

func (mu mu) unexportMethod() {
	fmt.Println("call unexport method from mu")
}

func TestIunexportMethod(t *testing.T) {
	muo := mu{}
	muo.unexportMethod()

	defer func() {
		r := recover()
		if r == nil || fmt.Sprintf("%v", r) != "runtime error: invalid memory address or nil pointer dereference" {
			t.Error(r)
		}
	}()
	// cannot use (mu literal) (value of type mu) as interfacex.IunexportMethod value in struct literal: mu does not implement interfacex.IunexportMethod (missing method unexportMethod)
	// 外部不可实现的接口，也就意味着这个接口只能被其所在包里的类型实现，可以确保实现的自主性，防止调用方使用自己的实现
	//
	// A sealed interface is an interface with unexported methods. This means users outside the package is unable to create types that fulfil the interface. This is useful for emulating a sum type as an exhaustive search for the types that fulfil the interface can be done.
	//
	// Read https://blog.chewxy.com/2018/03/18/golang-interfaces/#:~:text=A%20sealed%20interface%20is%20an%20interface%20with%20unexported,types%20that%20fulfil%20the%20interface%20can%20be%20done.
	m := MU{
		// IunexportMethod: mu{},
	}
	m.Method()

	// 标准库go/ast里的接口：
	var _ ast.Expr
	// 其定义：
	// All expression nodes implement the Expr interface.
	// type Expr interface {
	// 	Node
	// 	exprNode()
	// }
	// 其中`exprNode()`就是一个非导出方法
	var _ ast.IndexExpr
	// func (*IndexExpr) exprNode()      {}
}
