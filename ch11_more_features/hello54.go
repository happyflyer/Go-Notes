package main

import (
	"fmt"
)

// Phone Phone
type Phone interface {
	call()
}

// NokiaPhone NokiaPhone
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

// IPhone IPhone
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
