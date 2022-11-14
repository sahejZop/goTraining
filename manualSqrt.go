package main

import (
	"fmt"
)

func testFun() {
	fmt.Println("hi from file 2")
	testSwitchCase()
}

func deferReturn() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	defer fmt.Println("lol")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func testSwitchCase() {
	var a = 10
	switch a {
	case 5:
		fmt.Println(5)

	case 10:
		fmt.Println(10)
	}
}
