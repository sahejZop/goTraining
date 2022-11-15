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

func sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
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
