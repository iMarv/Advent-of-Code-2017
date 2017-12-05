package dayfive

type stepFunction func(instructions []int, index, step int) int

// CountSteps counts steps to take to exit
func CountSteps(instructions []int, step stepFunction) (steps int) {
	return step(instructions, 0, 0)
}

// StepForward steps forward always increasing offset
func StepForward(instructions []int, index, step int) int {
	if index >= 0 && index < len(instructions) {
		indexChange := instructions[index]
		instructions[index]++
		return StepForward(instructions, index+indexChange, step+1)
	}
	return step
}

// StepForwardDecreaseIterative solves second puzzle iteratively instead of recursively
func StepForwardDecreaseIterative(instructions []int) (steps int) {
	for index := 0; index < len(instructions); steps++ {
		indexChange := instructions[index]
		if indexChange > 2 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		index += indexChange
	}
	return
}

// StepForwardDecrease steps forward, decreasing offset if 3 or more
func StepForwardDecrease(instructions []int, index, step int) int {
	if index < len(instructions) {
		indexChange := instructions[index]
		if indexChange > 2 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		return StepForwardDecrease(instructions, index+indexChange, step+1)
	}
	return step
}
