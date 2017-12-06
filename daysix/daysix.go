package daysix

import (
	"fmt"
	"strings"
)

// CountRedistributions counts how many distributions are executed until a slice appears twice
func CountRedistributions(row []int) (steps int) {
	options := make(map[string]bool)
	options[arrayToString(row, ".")] = true

	for {
		RedistributeRow(row)
		steps++
		asString := arrayToString(row, ".")
		if options[asString] {
			break
		}
		options[asString] = true
	}

	return
}

// CountRedistributionsSingle counts the cycles until the start condition returns
func CountRedistributionsSingle(row []int) int {
	steps := 0
	options := make(map[string]int)
	options[arrayToString(row, ".")] = steps

	for {
		RedistributeRow(row)
		steps++
		asString := arrayToString(row, ".")
		if _, ok := options[asString]; ok {
			return steps - options[asString]
		}
		options[asString] = steps
	}
}

// RedistributeRow redistributes numbers in given row
func RedistributeRow(row []int) {
	highest := getIndexHighestNumber(row)
	num := row[highest]
	length := len(row)
	row[highest] = 0

	for index := range row {
		row[index] += num / length
	}
	rest := num % length

	for i := (highest + 1) % length; rest > 0; rest-- {
		row[i]++
		i = (i + 1) % length
	}
}

func getIndexHighestNumber(numbers []int) int {
	highest := 0
	index := 0
	for i, val := range numbers {
		if val > highest {
			index = i
			highest = val
		}
	}
	return index
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
