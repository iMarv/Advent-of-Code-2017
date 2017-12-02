package daytwo

import (
	"testing"
)

func TestCalcFirstChecksum(t *testing.T) {
	values := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	actual := CalcFirstChecksum(values)
	expected := 18
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCalcSecondChecksum(t *testing.T) {
	values := [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}
	actual := CalcSecondChecksum(values)
	expected := 9
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestMinAndMax(t *testing.T) {
	values := []int{5, 1, 9, 5}
	actualMin, actualMax := minAndMax(values)
	expectedMin := 1
	expectedMax := 9

	if actualMin != expectedMin {
		t.Fatalf("Expected min to be %d but got %d", expectedMin, actualMin)
	}

	if actualMax != expectedMax {
		t.Fatalf("Expected min to be %d but got %d", expectedMax, actualMax)
	}
}

func TestDivideCheckum5928(t *testing.T) {
	values := []int{5, 9, 2, 8}
	actual := divideChecksum(values)
	expected := 4

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestDivideCheckum9473(t *testing.T) {
	values := []int{9, 4, 7, 3}
	actual := divideChecksum(values)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestDivideCheckum3865(t *testing.T) {
	values := []int{3, 8, 6, 5}
	actual := divideChecksum(values)
	expected := 2

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
