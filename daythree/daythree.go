package daythree

import "math"

func CalcSteps(num int) (steps int) {
	return
}

func rowLength(num int) (rowLength int) {
	rowLength = int(math.Ceil(math.Sqrt(float64(num))))
	if rowLength%2 == 0 {
		rowLength += 1
	}
	return
}

func circleStep(num int) int {
	return (num - 1) / 2
}
