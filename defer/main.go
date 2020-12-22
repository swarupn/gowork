package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello, playground")
	x1 := func() {
		fmt.Println("hi")
	}
	x1()
	m := 10
	defer foo(m)
	defer bar(m)
}

func foo(x int) {
	fmt.Println("Hello, foo ", x)

}

func bar(x int) {
	fmt.Println("Hello, bar ", x)

}
