package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	sum := 0
	for {
		sum++ // 无限循环下去
	}
	// fmt.Println(sum) // 无法输出

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}
