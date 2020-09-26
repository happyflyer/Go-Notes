package main

import (
	"fmt"
)

func main() {
	// a := 1.69
	// b := 1.7
	// c := a * b     // 结果应该是2.873
	// fmt.Println(c) // 输出的是2.8729999999999998

	a := 1690                         // 表示1.69
	b := 1700                         // 表示1.70
	c := a * b                        // 结果应该是2873000表示 2.873
	fmt.Println(c)                    // 内部编码
	fmt.Println(float64(c) / 1000000) // 显示
}
