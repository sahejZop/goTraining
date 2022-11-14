package main

import (
	"fmt"
	"math"
	"rsc.io/quote"
)

func multiply(x int, y int) int {
	return x * y
}

func swapStrings(str1 string, str2 string) (string, string) {
	return str2, str1
}

func square(n int) int {
	fmt.Println(quote.Go())
	return n * n
}

func greaterThan(a, b int) bool {
	return a > b
}

func sumOfSquares(squareTill int) int {
	sum := 0
	for i := 1; i <= squareTill; i++ {
		sum += square(i)
	}
	return sum
}

const testConst = 100

func printStuff() {
	intFromFloat := math.Sqrt(64)
	fmt.Printf("intFromFloat: %g\n", intFromFloat)
	fmt.Printf("multiplied val: %d\n", multiply(int(intFromFloat), 5))
	var str1 = "String 1"
	var str2 = "String 2"
	str1, str2 = swapStrings(str1, str2)
	i := 100
	fmt.Println(i)
	i++
	fmt.Println(i)
	//strI := i
	// ascii
	strI := string(i)
	fmt.Println(strI)
	fmt.Printf("str1 after swap: %s\n", str1)
	fmt.Printf("str2 after swap: %s\n", str2)
	fmt.Printf("Sum of squares: %v\n", sumOfSquares(9))
	fmt.Printf("4 in base 2: %b\n", 4)

	fmt.Println()

	var firstName string
	fmt.Println("Enter your first name: ")
	fmt.Scanf("%s", &firstName)
	fmt.Println("\nHi,", firstName)

	var (
		a int
		b int
	)

	fmt.Printf("Enter first number: \n")
	fmt.Scanf("%d", &a)
	fmt.Printf("Enter second number: \n")
	fmt.Scanf("%d", &b)

	if greaterThan(a, b) {
		fmt.Printf("a is greater than b\n")
	} else {
		fmt.Printf("b is greater than a\n")
	}

	testFun()
}

func main() {
	//printStuff()
	fmt.Println(quote.Go())
	//printStuff()
	deferReturn()
	fmt.Println(quote.Glass())
}
