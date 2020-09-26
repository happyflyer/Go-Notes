package main

import "fmt"

func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	// 声明一个变量并初始化
	var a = "RUNOOB"
	fmt.Println(a)
	// 没有初始化就为零值
	var b int
	fmt.Println(b)
	// bool 零值为 false
	var c bool
	fmt.Println(c)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 第二种，根据值自行判定变量类型。
	var d = true
	fmt.Println(d)

	// 第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)
}
