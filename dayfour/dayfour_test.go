package dayfour

import (
	"testing"
)

func TestValidPhrasesCountDuplicates(t *testing.T) {
	passphraseCollection := [][]string{
		{"aa", "bb", "cc", "dd", "ee"},
		{"aa", "bb", "cc", "dd", "aa"},
		{"aa", "bb", "cc", "dd", "aaa"},
	}

	actual := ValidPhrasesCount(passphraseCollection, NoDuplicatesValidator)
	expected := 2

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestValidPhrasesCountAnagram(t *testing.T) {
	passphraseCollection := [][]string{
		{"abcde", "fghij"},
		{"abcde", "xyz", "ecdab"},
		{"a", "ab", "abc", "abd", "abf", "abj"},
		{"iiii", "oiii", "ooii", "oooi", "oooo"},
		{"oiii", "ioii", "iioi", "iiio"},
	}

	actual := ValidPhrasesCount(passphraseCollection, NoAnagramsValidator)
	expected := 3

	if actual != expected {
		t.Fatalf("Expected %d but got %d", expected, actual)
	}
}

func TestNoDuplicatesValidator1(t *testing.T) {
	phrases := []string{"aa", "bb", "cc", "dd", "ee"}
	actual := NoDuplicatesValidator(phrases)
	expected := true

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoDuplicatesValidator2(t *testing.T) {
	phrases := []string{"aa", "bb", "cc", "dd", "aa"}
	actual := NoDuplicatesValidator(phrases)
	expected := false

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoDuplicatesValidator3(t *testing.T) {
	phrases := []string{"aa", "bb", "cc", "dd", "aaa"}
	actual := NoDuplicatesValidator(phrases)
	expected := true

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoAnagramsValidator1(t *testing.T) {
	phrases := []string{"abcde", "fghij"}
	actual := NoAnagramsValidator(phrases)
	expected := true

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoAnagramsValidator2(t *testing.T) {
	phrases := []string{"abcde", "xyz", "ecdab"}
	actual := NoAnagramsValidator(phrases)
	expected := false

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoAnagramsValidator3(t *testing.T) {
	phrases := []string{"a", "ab", "abc", "abd", "abf", "abj"}
	actual := NoAnagramsValidator(phrases)
	expected := true

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoAnagramsValidator4(t *testing.T) {
	phrases := []string{"iiii", "oiii", "ooii", "oooi", "oooo"}
	actual := NoAnagramsValidator(phrases)
	expected := true

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestNoAnagramsValidator5(t *testing.T) {
	phrases := []string{"oiii", "ioii", "iioi", "iiio"}
	actual := NoAnagramsValidator(phrases)
	expected := false

	if actual != expected {
		t.Fatalf("Expected %t but got %t", expected, actual)
	}
}

func TestOrderString(t *testing.T) {
	actual := orderString("dcba")
	expected := "abcd"

	if actual != expected {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}
