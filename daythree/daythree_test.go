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

func TestRowLength1(t *testing.T) {
	actual := rowLength(1)
	expected := 1

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestRowLength8(t *testing.T) {
	actual := rowLength(8)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestRowLength23(t *testing.T) {
	actual := rowLength(23)
	expected := 5

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestRowLength27(t *testing.T) {
	actual := rowLength(27)
	expected := 7

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCircleStep1(t *testing.T) {
	actual := circleStep(1)
	expected := 0

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCircleStep3(t *testing.T) {
	actual := circleStep(3)
	expected := 1

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCircleStep5(t *testing.T) {
	actual := circleStep(5)
	expected := 2

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCircleStep7(t *testing.T) {
	actual := circleStep(7)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
