.PHONY:

# 因为在go module里面，而目前的版本还是1.17，因此使用gotip执行时关闭module，否则报错：
# github.com/donnol/blog/demo/go/generic
# ./append.go:4:2: embedding interface element string|[]byte requires go1.18 or later (-lang was set to go1.17; check go.mod)
# ./append.go:8:23: type parameters require go1.18 or later
build:
	GO111MODULE=off gotip build
