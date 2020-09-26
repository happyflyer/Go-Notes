package main

import "fmt"

const (
	i = 1 << iota // 1 << 0
	j = 3 << iota // 11 << 1 == 110 == 6
	k             // 11 << 2 == 1100 == 12
	l             // 11 << 3 == 11000 == 24
	// << n == *(2^n)
)

func main() {
	fmt.Println("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)
	fmt.Println("l =", l)
}
