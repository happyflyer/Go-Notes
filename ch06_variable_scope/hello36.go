package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
	/* 声明局部变量 */
	var g int = 10

	fmt.Printf("结果： g = %d\n", g)
}
