package dayone

import (
	"testing"
)

func TestFirstCaptcha1122(t *testing.T) {
	actual := FirstCaptcha("1122")
	expected := 3
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestFirstCaptcha1111(t *testing.T) {
	actual := FirstCaptcha("1111")
	expected := 4
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestFirstCaptcha1234(t *testing.T) {
	actual := FirstCaptcha("1234")
	expected := 0
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestFirstCaptcha91212129(t *testing.T) {
	actual := FirstCaptcha("91212129")
	expected := 9
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSecondCaptcha1212(t *testing.T) {
	actual := SecondCaptcha("1212")
	expected := 6
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSecondCaptcha1221(t *testing.T) {
	actual := SecondCaptcha("1221")
	expected := 0
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSecondCaptcha123425(t *testing.T) {
	actual := SecondCaptcha("123425")
	expected := 4
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSecondCaptcha123123(t *testing.T) {
	actual := SecondCaptcha("123123")
	expected := 12
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestSecondCaptcha12131415(t *testing.T) {
	actual := SecondCaptcha("12131415")
	expected := 4
	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}
