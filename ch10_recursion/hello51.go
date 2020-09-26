package main

import "fmt"

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, factorial(uint64(i)))
}
