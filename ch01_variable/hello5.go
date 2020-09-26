package main

import "unsafe"

// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a, b, c)
}
