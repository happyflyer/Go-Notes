package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello"

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
	println(&g, &h)

	// 如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明

	// 如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
	var a string = "abc"
	fmt.Println("hello, world", a)
}
