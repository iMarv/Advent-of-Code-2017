package daythree

import (
	"testing"
)

func TestCalcSteps1(t *testing.T) {
	actual := CalcSteps(1)
	expected := 0

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCalcSteps12(t *testing.T) {
	actual := CalcSteps(12)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCalcSteps23(t *testing.T) {
	actual := CalcSteps(23)
	expected := 2

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCalcSteps1024(t *testing.T) {
	actual := CalcSteps(1024)
	expected := 31

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
