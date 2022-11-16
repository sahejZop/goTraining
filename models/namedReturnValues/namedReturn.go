package namedReturnValues

import "math"

func RoundOff(a, b float64) (x, y float64) {
	x = math.Floor(a + 10)
	y = math.Floor(b + 5)
	defer func() {
		x *= 10
		y *= 10
	}()
	return
}
