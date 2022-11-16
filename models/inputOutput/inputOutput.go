package inputOutput

import (
	"fmt"
)

const (
	x int = 10
	y
	z
)

func InputOutput() {
	var var1 int
	var1, _ = fmt.Printf("Hello World")
	fmt.Println(var1)
}

func InputOutput2() {
	fmt.Println(x, y, z)
}

//func InputOutput3() {
//	var var1 float64 = 10.35
//	var var2 float32 = 10.35
//
//	// gives error
//	//if var1 == var2 {
//	//
//	//}
//}

func InputOutput4() {
	var var1 float32 = 10.35

	if var1 == 10.35 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Hiiii")
	}
}

func InputOutput5() {
	var var1 bool = true
	if var1 == true {
		fmt.Println("hiii")
	}
}

func InputOutput6() {
	var var1 string
	var1 = "Hello World"

	if var1 == "Hello World" {
		fmt.Println("Hello")
	} else {
		fmt.Println("Hiii")
	}
}
