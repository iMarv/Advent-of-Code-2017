package daysix

import (
	"reflect"
	"testing"
)

func TestCountRedistributions0270(t *testing.T) {
	vals := []int{0, 2, 7, 0}
	actual := CountRedistributions(vals)
	expected := 5

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCountRedistributionsSingle2412(t *testing.T) {
	vals := []int{2, 4, 1, 2}
	actual := CountRedistributionsSingle(vals)
	expected := 4

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestRedistributeRow0270(t *testing.T) {
	vals := []int{0, 2, 7, 0}
	RedistributeRow(vals)
	expected := []int{2, 4, 1, 2}

	if !reflect.DeepEqual(vals, expected) {
		t.Fatalf("Expected %s but got %s", arrayToString(expected, ""), arrayToString(vals, ""))
	}
}

func TestRedistributeRow2412(t *testing.T) {
	vals := []int{2, 4, 1, 2}
	RedistributeRow(vals)
	expected := []int{3, 1, 2, 3}

	if !reflect.DeepEqual(vals, expected) {
		t.Fatalf("Expected %s but got %s", arrayToString(expected, ""), arrayToString(vals, ""))
	}
}

func TestRedistributeRow3123(t *testing.T) {
	vals := []int{3, 1, 2, 3}
	RedistributeRow(vals)
	expected := []int{0, 2, 3, 4}

	if !reflect.DeepEqual(vals, expected) {
		t.Fatalf("Expected %s but got %s", arrayToString(expected, ""), arrayToString(vals, ""))
	}
}

func TestRedistributeRow0234(t *testing.T) {
	vals := []int{0, 2, 3, 4}
	RedistributeRow(vals)
	expected := []int{1, 3, 4, 1}

	if !reflect.DeepEqual(vals, expected) {
		t.Fatalf("Expected %s but got %s", arrayToString(expected, ""), arrayToString(vals, ""))
	}
}

func TestGetIndexHighestNumber0234(t *testing.T) {
	vals := []int{0, 2, 3, 4}
	actual := getIndexHighestNumber(vals)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestGetIndexHighestNumber7234(t *testing.T) {
	vals := []int{7, 2, 3, 4}
	actual := getIndexHighestNumber(vals)
	expected := 0

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
