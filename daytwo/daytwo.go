package daytwo

import "sort"

// CalcFirstChecksum calculates the checksum of a two dimensional slice
func CalcFirstChecksum(values [][]int) (checksum int) {
	for _, row := range values {
		min, max := minAndMax(row)
		checksum += max - min
	}
	return
}

// CalcSecondChecksum calculates the checksum of a two dimensional slice
func CalcSecondChecksum(values [][]int) (checksum int) {
	for _, row := range values {
		checksum += divideChecksum(row)
	}
	return
}

func minAndMax(numbers []int) (int, int) {
	var min, max int
	min = numbers[0]
	for _, val := range numbers {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	return min, max
}

func divideChecksum(numbers []int) (sum int) {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	for big := 0; big < len(numbers); big++ {
		if big < len(numbers)-1 {
			for small := big + 1; small < len(numbers); small++ {
				if numbers[big]%numbers[small] == 0 {
					sum += numbers[big] / numbers[small]
				}
			}
		}
	}
	return
}
