package main

import (
	"fmt"
	"github.com/sahejZop/goTraining/structs"
	"math"
)

func testPointers() {
	var p *int
	var a = 100
	p = &a
	fmt.Println(*p)
}

func DistanceOfTwoCartesianPoints(a, b structs.CartesianPoint) float64 {
	return math.Sqrt(float64(square(a.X-b.X) + square(a.Y-b.Y)))
}
