package main

import "fmt"

func sliceOperations() {
	sliceList := make([]int, 0, 3)
	fmt.Println(sliceList)
	for i := 0; i < 5; i++ {
		sliceList = append(sliceList, i)
		fmt.Printf("capacity: %v\nlen: %v\nsliceList: %p\n", cap(sliceList), len(sliceList), sliceList)
	}
}

func moreSliceOperations() {

}
