package dayfive

import "testing"

func TestCountSteps(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}

	actual := CountSteps(instructions, StepForward)
	expected := 5

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCountStepsDecrease(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}

	actual := CountSteps(instructions, StepForwardDecrease)
	expected := 10

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestCountStepsDecreaseIterative(t *testing.T) {
	instructions := []int{0, 3, 0, 1, -3}

	actual := StepForwardDecreaseIterative(instructions)
	expected := 10

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func BenchmarkCountStepsDecrease100(b *testing.B) {

	for n := 0; n < b.N; n++ {
		instructions := []int{0, 3, 0, 1, -3}
		CountSteps(instructions, StepForwardDecrease)

	}
}
